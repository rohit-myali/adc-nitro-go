/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package cs

/**
* Binding class showing the appflowpolicy that can be bound to csvserver.
*/
type Csvserverappflowpolicybinding struct {
	/**
	* Policies bound to this vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority for the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke flag.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* The invocation type.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label invoked.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
		Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1
		Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver
	*/
	Targetlbvserver string `json:"targetlbvserver,omitempty"`
	/**
	* Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`


}