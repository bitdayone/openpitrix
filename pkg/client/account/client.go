// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package account

import (
	"context"
	"fmt"
	"math"
	"strings"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/stringutil"
)

type Client struct {
	pb.AccountManagerClient
}

func NewClient() (*Client, error) {
	conn, err := manager.NewClient(constants.AccountServiceHost, constants.AccountServicePort)
	if err != nil {
		return nil, err
	}
	return &Client{
		AccountManagerClient: pb.NewAccountManagerClient(conn),
	}, nil
}

func (c *Client) GetUsers(ctx context.Context, userIds []string) ([]*pb.User, error) {
	var internalUsers []*pb.User
	var noInternalUserIds []string
	for _, userId := range userIds {
		if stringutil.StringIn(userId, constants.InternalUsers) {
			internalUsers = append(internalUsers, &pb.User{
				UserId: pbutil.ToProtoString(userId),
			})
		} else {
			noInternalUserIds = append(noInternalUserIds, userId)
		}
	}

	if len(noInternalUserIds) == 0 {
		return internalUsers, nil
	}

	response, err := c.DescribeUsers(ctx, &pb.DescribeUsersRequest{
		UserId: noInternalUserIds,
	})
	if err != nil {
		logger.Error(ctx, "Describe users %s failed: %+v", noInternalUserIds, err)
		return nil, err
	}
	if len(response.UserSet) != len(noInternalUserIds) {
		logger.Error(ctx, "Describe users %s with return count [%d]", userIds, len(response.UserSet)+len(internalUsers))
		return nil, fmt.Errorf("describe users %s with return count [%d]", userIds, len(response.UserSet)+len(internalUsers))
	}
	response.UserSet = append(response.UserSet, internalUsers...)
	return response.UserSet, nil
}

func (c *Client) GetUserGroupPath(ctx context.Context, userId string) (string, error) {
	var userGroupPath string

	response, err := c.DescribeUsersDetail(ctx, &pb.DescribeUsersRequest{
		UserId: []string{userId},
	})
	if err != nil || len(response.UserDetailSet) == 0 {
		logger.Error(ctx, "Describe user [%s] failed: %+v", userId, err)
		return "", err
	}

	groups := response.UserDetailSet[0].GroupSet

	//If one user under different groups, get the highest group path.
	minLevel := math.MaxInt32
	for _, group := range groups {
		level := len(strings.Split(group.GroupPath.GetValue(), "."))
		if level < minLevel {
			minLevel = level
			userGroupPath = group.GetGroupPath().GetValue()
		}
	}

	return userGroupPath, nil

}

func GetRoleUsers(ctx context.Context, roleIds []string) ([]*pb.User, error) {
	client, err := NewClient()
	if err != nil {
		logger.Error(ctx, "Failed to create account manager client: %+v", err)
		return nil, err
	}

	response, err := client.DescribeUsers(ctx, &pb.DescribeUsersRequest{
		RoleId: roleIds,
	})
	if err != nil {
		logger.Error(ctx, "Describe users failed: %+v", err)
		return nil, err
	}

	return response.UserSet, nil
}

func GetUsers(ctx context.Context, userIds []string) ([]*pb.User, error) {
	client, err := NewClient()
	if err != nil {
		logger.Error(ctx, "Failed to create account manager client: %+v", err)
		return nil, err
	}
	response, err := client.GetUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	return response, err
}

func GetIsvFromUsers(ctx context.Context, userIds []string) ([]*pb.User, error) {
	client, err := NewClient()
	if err != nil {
		logger.Error(ctx, "Failed to create account manager client: %+v", err)
		return nil, err
	}

	var owners []string
	for _, userId := range userIds {
		response, err := client.GetUserGroupOwner(ctx, &pb.GetUserGroupOwnerRequest{
			UserId: userId,
		})
		if err != nil {
			return nil, err
		}
		owners = append(owners, response.Owner)
	}

	response, err := client.GetUsers(ctx, owners)
	if err != nil {
		return nil, err
	}
	return response, err
}
