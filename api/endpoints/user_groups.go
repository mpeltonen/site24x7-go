package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type UserGroupsEndpoint interface {
	Get(groupID string) (*api.UserGroup, error)
	Create(group *api.UserGroup) (*api.UserGroup, error)
	Update(group *api.UserGroup) (*api.UserGroup, error)
	Delete(groupID string) error
	List() ([]*api.UserGroup, error)
}

type userGroupsEndpoint struct {
	client rest.Client
}

func NewUserGroupsEndpoint(client rest.Client) UserGroupsEndpoint {
	return &userGroupsEndpoint{
		client: client,
	}
}

func (c *userGroupsEndpoint) Get(groupID string) (*api.UserGroup, error) {
	userGroup := &api.UserGroup{}
	err := c.client.
		Get().
		Resource("user_groups").
		ResourceID(groupID).
		Do().
		Into(userGroup)

	return userGroup, err
}

func (c *userGroupsEndpoint) Create(group *api.UserGroup) (*api.UserGroup, error) {
	newUserGroup := &api.UserGroup{}
	err := c.client.
		Post().
		Resource("user_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Into(newUserGroup)

	return newUserGroup, err
}

func (c *userGroupsEndpoint) Update(group *api.UserGroup) (*api.UserGroup, error) {
	updatedGroup := &api.UserGroup{}
	err := c.client.
		Put().
		Resource("user_groups").
		ResourceID(group.UserGroupID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Into(updatedGroup)

	return updatedGroup, err
}

func (c *userGroupsEndpoint) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("user_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *userGroupsEndpoint) List() ([]*api.UserGroup, error) {
	userGroups := []*api.UserGroup{}
	err := c.client.
		Get().
		Resource("user_groups").
		Do().
		Into(&userGroups)

	return userGroups, err
}
