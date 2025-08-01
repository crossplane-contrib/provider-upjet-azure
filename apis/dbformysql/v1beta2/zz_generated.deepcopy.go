//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyInitParameters) DeepCopyInto(out *CustomerManagedKeyInitParameters) {
	*out = *in
	if in.GeoBackupKeyVaultKeyID != nil {
		in, out := &in.GeoBackupKeyVaultKeyID, &out.GeoBackupKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.GeoBackupUserAssignedIdentityID != nil {
		in, out := &in.GeoBackupUserAssignedIdentityID, &out.GeoBackupUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityID != nil {
		in, out := &in.PrimaryUserAssignedIdentityID, &out.PrimaryUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyInitParameters.
func (in *CustomerManagedKeyInitParameters) DeepCopy() *CustomerManagedKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyObservation) DeepCopyInto(out *CustomerManagedKeyObservation) {
	*out = *in
	if in.GeoBackupKeyVaultKeyID != nil {
		in, out := &in.GeoBackupKeyVaultKeyID, &out.GeoBackupKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.GeoBackupUserAssignedIdentityID != nil {
		in, out := &in.GeoBackupUserAssignedIdentityID, &out.GeoBackupUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityID != nil {
		in, out := &in.PrimaryUserAssignedIdentityID, &out.PrimaryUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyObservation.
func (in *CustomerManagedKeyObservation) DeepCopy() *CustomerManagedKeyObservation {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerManagedKeyParameters) DeepCopyInto(out *CustomerManagedKeyParameters) {
	*out = *in
	if in.GeoBackupKeyVaultKeyID != nil {
		in, out := &in.GeoBackupKeyVaultKeyID, &out.GeoBackupKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.GeoBackupUserAssignedIdentityID != nil {
		in, out := &in.GeoBackupUserAssignedIdentityID, &out.GeoBackupUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityID != nil {
		in, out := &in.PrimaryUserAssignedIdentityID, &out.PrimaryUserAssignedIdentityID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerManagedKeyParameters.
func (in *CustomerManagedKeyParameters) DeepCopy() *CustomerManagedKeyParameters {
	if in == nil {
		return nil
	}
	out := new(CustomerManagedKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServer) DeepCopyInto(out *FlexibleServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServer.
func (in *FlexibleServer) DeepCopy() *FlexibleServer {
	if in == nil {
		return nil
	}
	out := new(FlexibleServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerInitParameters) DeepCopyInto(out *FlexibleServerInitParameters) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordSecretRef != nil {
		in, out := &in.AdministratorPasswordSecretRef, &out.AdministratorPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AdministratorPasswordWo != nil {
		in, out := &in.AdministratorPasswordWo, &out.AdministratorPasswordWo
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordWoVersion != nil {
		in, out := &in.AdministratorPasswordWoVersion, &out.AdministratorPasswordWoVersion
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(CustomerManagedKeyInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegatedSubnetID != nil {
		in, out := &in.DelegatedSubnetID, &out.DelegatedSubnetID
		*out = new(string)
		**out = **in
	}
	if in.DelegatedSubnetIDRef != nil {
		in, out := &in.DelegatedSubnetIDRef, &out.DelegatedSubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegatedSubnetIDSelector != nil {
		in, out := &in.DelegatedSubnetIDSelector, &out.DelegatedSubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GeoRedundantBackupEnabled != nil {
		in, out := &in.GeoRedundantBackupEnabled, &out.GeoRedundantBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailabilityInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PointInTimeRestoreTimeInUtc != nil {
		in, out := &in.PointInTimeRestoreTimeInUtc, &out.PointInTimeRestoreTimeInUtc
		*out = new(string)
		**out = **in
	}
	if in.PrivateDNSZoneID != nil {
		in, out := &in.PrivateDNSZoneID, &out.PrivateDNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.PrivateDNSZoneIDRef != nil {
		in, out := &in.PrivateDNSZoneIDRef, &out.PrivateDNSZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateDNSZoneIDSelector != nil {
		in, out := &in.PrivateDNSZoneIDSelector, &out.PrivateDNSZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SourceServerID != nil {
		in, out := &in.SourceServerID, &out.SourceServerID
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerInitParameters.
func (in *FlexibleServerInitParameters) DeepCopy() *FlexibleServerInitParameters {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerList) DeepCopyInto(out *FlexibleServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerList.
func (in *FlexibleServerList) DeepCopy() *FlexibleServerList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerObservation) DeepCopyInto(out *FlexibleServerObservation) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordWo != nil {
		in, out := &in.AdministratorPasswordWo, &out.AdministratorPasswordWo
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordWoVersion != nil {
		in, out := &in.AdministratorPasswordWoVersion, &out.AdministratorPasswordWoVersion
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(CustomerManagedKeyObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegatedSubnetID != nil {
		in, out := &in.DelegatedSubnetID, &out.DelegatedSubnetID
		*out = new(string)
		**out = **in
	}
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
	if in.GeoRedundantBackupEnabled != nil {
		in, out := &in.GeoRedundantBackupEnabled, &out.GeoRedundantBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailabilityObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.PointInTimeRestoreTimeInUtc != nil {
		in, out := &in.PointInTimeRestoreTimeInUtc, &out.PointInTimeRestoreTimeInUtc
		*out = new(string)
		**out = **in
	}
	if in.PrivateDNSZoneID != nil {
		in, out := &in.PrivateDNSZoneID, &out.PrivateDNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ReplicaCapacity != nil {
		in, out := &in.ReplicaCapacity, &out.ReplicaCapacity
		*out = new(float64)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SourceServerID != nil {
		in, out := &in.SourceServerID, &out.SourceServerID
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerObservation.
func (in *FlexibleServerObservation) DeepCopy() *FlexibleServerObservation {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerParameters) DeepCopyInto(out *FlexibleServerParameters) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordSecretRef != nil {
		in, out := &in.AdministratorPasswordSecretRef, &out.AdministratorPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AdministratorPasswordWo != nil {
		in, out := &in.AdministratorPasswordWo, &out.AdministratorPasswordWo
		*out = new(string)
		**out = **in
	}
	if in.AdministratorPasswordWoVersion != nil {
		in, out := &in.AdministratorPasswordWoVersion, &out.AdministratorPasswordWoVersion
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKey != nil {
		in, out := &in.CustomerManagedKey, &out.CustomerManagedKey
		*out = new(CustomerManagedKeyParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegatedSubnetID != nil {
		in, out := &in.DelegatedSubnetID, &out.DelegatedSubnetID
		*out = new(string)
		**out = **in
	}
	if in.DelegatedSubnetIDRef != nil {
		in, out := &in.DelegatedSubnetIDRef, &out.DelegatedSubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegatedSubnetIDSelector != nil {
		in, out := &in.DelegatedSubnetIDSelector, &out.DelegatedSubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GeoRedundantBackupEnabled != nil {
		in, out := &in.GeoRedundantBackupEnabled, &out.GeoRedundantBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailabilityParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PointInTimeRestoreTimeInUtc != nil {
		in, out := &in.PointInTimeRestoreTimeInUtc, &out.PointInTimeRestoreTimeInUtc
		*out = new(string)
		**out = **in
	}
	if in.PrivateDNSZoneID != nil {
		in, out := &in.PrivateDNSZoneID, &out.PrivateDNSZoneID
		*out = new(string)
		**out = **in
	}
	if in.PrivateDNSZoneIDRef != nil {
		in, out := &in.PrivateDNSZoneIDRef, &out.PrivateDNSZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateDNSZoneIDSelector != nil {
		in, out := &in.PrivateDNSZoneIDSelector, &out.PrivateDNSZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SourceServerID != nil {
		in, out := &in.SourceServerID, &out.SourceServerID
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerParameters.
func (in *FlexibleServerParameters) DeepCopy() *FlexibleServerParameters {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerSpec) DeepCopyInto(out *FlexibleServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerSpec.
func (in *FlexibleServerSpec) DeepCopy() *FlexibleServerSpec {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerStatus) DeepCopyInto(out *FlexibleServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerStatus.
func (in *FlexibleServerStatus) DeepCopy() *FlexibleServerStatus {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailabilityInitParameters) DeepCopyInto(out *HighAvailabilityInitParameters) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailabilityInitParameters.
func (in *HighAvailabilityInitParameters) DeepCopy() *HighAvailabilityInitParameters {
	if in == nil {
		return nil
	}
	out := new(HighAvailabilityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailabilityObservation) DeepCopyInto(out *HighAvailabilityObservation) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailabilityObservation.
func (in *HighAvailabilityObservation) DeepCopy() *HighAvailabilityObservation {
	if in == nil {
		return nil
	}
	out := new(HighAvailabilityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailabilityParameters) DeepCopyInto(out *HighAvailabilityParameters) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailabilityParameters.
func (in *HighAvailabilityParameters) DeepCopy() *HighAvailabilityParameters {
	if in == nil {
		return nil
	}
	out := new(HighAvailabilityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityInitParameters) DeepCopyInto(out *IdentityInitParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityInitParameters.
func (in *IdentityInitParameters) DeepCopy() *IdentityInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowInitParameters) DeepCopyInto(out *MaintenanceWindowInitParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowInitParameters.
func (in *MaintenanceWindowInitParameters) DeepCopy() *MaintenanceWindowInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowObservation) DeepCopyInto(out *MaintenanceWindowObservation) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowObservation.
func (in *MaintenanceWindowObservation) DeepCopy() *MaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowParameters) DeepCopyInto(out *MaintenanceWindowParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowParameters.
func (in *MaintenanceWindowParameters) DeepCopy() *MaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageInitParameters) DeepCopyInto(out *StorageInitParameters) {
	*out = *in
	if in.AutoGrowEnabled != nil {
		in, out := &in.AutoGrowEnabled, &out.AutoGrowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IoScalingEnabled != nil {
		in, out := &in.IoScalingEnabled, &out.IoScalingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.LogOnDiskEnabled != nil {
		in, out := &in.LogOnDiskEnabled, &out.LogOnDiskEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SizeGb != nil {
		in, out := &in.SizeGb, &out.SizeGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageInitParameters.
func (in *StorageInitParameters) DeepCopy() *StorageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
	if in.AutoGrowEnabled != nil {
		in, out := &in.AutoGrowEnabled, &out.AutoGrowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IoScalingEnabled != nil {
		in, out := &in.IoScalingEnabled, &out.IoScalingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.LogOnDiskEnabled != nil {
		in, out := &in.LogOnDiskEnabled, &out.LogOnDiskEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SizeGb != nil {
		in, out := &in.SizeGb, &out.SizeGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	if in.AutoGrowEnabled != nil {
		in, out := &in.AutoGrowEnabled, &out.AutoGrowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IoScalingEnabled != nil {
		in, out := &in.IoScalingEnabled, &out.IoScalingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.LogOnDiskEnabled != nil {
		in, out := &in.LogOnDiskEnabled, &out.LogOnDiskEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SizeGb != nil {
		in, out := &in.SizeGb, &out.SizeGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}
