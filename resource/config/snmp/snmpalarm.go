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

package snmp

/**
* Configuration for alarm resource.
*/
type Snmpalarm struct {
	/**
	* Name of the SNMP alarm. This parameter is required for identifying the SNMP alarm and cannot be modified.
	*/
	Trapname string `json:"trapname,omitempty"`
	/**
	* Value for the high threshold. The Citrix ADC generates an SNMP trap message when the value of the attribute associated with the alarm is greater than or equal to the specified high threshold value.
	*/
	Thresholdvalue int `json:"thresholdvalue,omitempty"`
	/**
	* Value for the normal threshold. A trap message is generated if the value of the respective attribute falls to or below this value after exceeding the high threshold.
	*/
	Normalvalue int `json:"normalvalue,omitempty"`
	/**
	* Interval, in seconds, at which the Citrix ADC generates SNMP trap messages when the conditions specified in the SNMP alarm are met.Can be specified for the following alarms: SYNFLOOD, HA-VERSION-MISMATCH, HA-SYNC-FAILURE, HA-NO-HEARTBEATS,HA-BAD-SECONDARY-STATE, CLUSTER-NODE-HEALTH, CLUSTER-NODE-QUORUM, CLUSTER-VERSION-MISMATCH, CLUSTER-BKHB-FAILED, PORT-ALLOC-FAILED, COMPACT-FLASH-ERRORS, HARD-DISK-DRIVE-ERRORS and APPFW traps. Default trap time intervals: SYNFLOOD and APPFW traps = 1sec, PORT-ALLOC-FAILED = 3600sec(1 hour), Other Traps = 86400sec(1 day)
	*/
	Time int `json:"time,omitempty"`
	/**
	* Current state of the SNMP alarm. The Citrix ADC generates trap messages only for SNMP alarms that are enabled. Some alarms are enabled by default, but you can disable them.
	*/
	State string `json:"state,omitempty"`
	/**
	* Severity level assigned to trap messages generated by this alarm. The severity levels are, in increasing order of severity, Informational, Warning, Minor, Major, and Critical.
		This parameter is useful when you want the Citrix ADC to send trap messages to a trap listener on the basis of severity level. Trap messages with a severity level lower than the specified level (in the trap listener entry) are not sent.
	*/
	Severity string `json:"severity,omitempty"`
	/**
	* Logging status of the alarm. When logging is enabled, the Citrix ADC logs every trap message that is generated for this alarm.
	*/
	Logging string `json:"logging,omitempty"`

	//------- Read only Parameter ---------;

	Timeout string `json:"timeout,omitempty"`

}
