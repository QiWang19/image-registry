// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Action, InType: reflect.TypeOf(&Action{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterPolicy, InType: reflect.TypeOf(&ClusterPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterPolicyBinding, InType: reflect.TypeOf(&ClusterPolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterPolicyBindingList, InType: reflect.TypeOf(&ClusterPolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterPolicyList, InType: reflect.TypeOf(&ClusterPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterRole, InType: reflect.TypeOf(&ClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterRoleBinding, InType: reflect.TypeOf(&ClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterRoleBindingList, InType: reflect.TypeOf(&ClusterRoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ClusterRoleList, InType: reflect.TypeOf(&ClusterRoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GroupRestriction, InType: reflect.TypeOf(&GroupRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_IsPersonalSubjectAccessReview, InType: reflect.TypeOf(&IsPersonalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LocalResourceAccessReview, InType: reflect.TypeOf(&LocalResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LocalSubjectAccessReview, InType: reflect.TypeOf(&LocalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamedClusterRole, InType: reflect.TypeOf(&NamedClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamedClusterRoleBinding, InType: reflect.TypeOf(&NamedClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamedRole, InType: reflect.TypeOf(&NamedRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamedRoleBinding, InType: reflect.TypeOf(&NamedRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PolicyBinding, InType: reflect.TypeOf(&PolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PolicyBindingList, InType: reflect.TypeOf(&PolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PolicyList, InType: reflect.TypeOf(&PolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ResourceAccessReview, InType: reflect.TypeOf(&ResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ResourceAccessReviewResponse, InType: reflect.TypeOf(&ResourceAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Role, InType: reflect.TypeOf(&Role{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleBinding, InType: reflect.TypeOf(&RoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleBindingList, InType: reflect.TypeOf(&RoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleBindingRestriction, InType: reflect.TypeOf(&RoleBindingRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleBindingRestrictionList, InType: reflect.TypeOf(&RoleBindingRestrictionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleBindingRestrictionSpec, InType: reflect.TypeOf(&RoleBindingRestrictionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoleList, InType: reflect.TypeOf(&RoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SelfSubjectRulesReview, InType: reflect.TypeOf(&SelfSubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SelfSubjectRulesReviewSpec, InType: reflect.TypeOf(&SelfSubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ServiceAccountReference, InType: reflect.TypeOf(&ServiceAccountReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ServiceAccountRestriction, InType: reflect.TypeOf(&ServiceAccountRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SubjectAccessReview, InType: reflect.TypeOf(&SubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SubjectAccessReviewResponse, InType: reflect.TypeOf(&SubjectAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SubjectRulesReview, InType: reflect.TypeOf(&SubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SubjectRulesReviewSpec, InType: reflect.TypeOf(&SubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SubjectRulesReviewStatus, InType: reflect.TypeOf(&SubjectRulesReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserRestriction, InType: reflect.TypeOf(&UserRestriction{})},
	)
}

// DeepCopy_v1_Action is an autogenerated deepcopy function.
func DeepCopy_v1_Action(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Action)
		out := out.(*Action)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Content); err != nil {
			return err
		} else {
			out.Content = *newVal.(*runtime.RawExtension)
		}
		return nil
	}
}

// DeepCopy_v1_ClusterPolicy is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicy)
		out := out.(*ClusterPolicy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(NamedClusterRoles, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_NamedClusterRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterPolicyBinding is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterPolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBinding)
		out := out.(*ClusterPolicyBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(NamedClusterRoleBindings, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_NamedClusterRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterPolicyBindingList is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterPolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBindingList)
		out := out.(*ClusterPolicyBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ClusterPolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterPolicyList is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyList)
		out := out.(*ClusterPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ClusterPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterRole is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRole)
		out := out.(*ClusterRole)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterRoleBinding is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBinding)
		out := out.(*ClusterRoleBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if in.UserNames != nil {
			in, out := &in.UserNames, &out.UserNames
			*out = make(OptionalNames, len(*in))
			copy(*out, *in)
		}
		if in.GroupNames != nil {
			in, out := &in.GroupNames, &out.GroupNames
			*out = make(OptionalNames, len(*in))
			copy(*out, *in)
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]api_v1.ObjectReference, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_ClusterRoleBindingList is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterRoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBindingList)
		out := out.(*ClusterRoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ClusterRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_ClusterRoleList is an autogenerated deepcopy function.
func DeepCopy_v1_ClusterRoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleList)
		out := out.(*ClusterRoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRole, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ClusterRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_GroupRestriction is an autogenerated deepcopy function.
func DeepCopy_v1_GroupRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupRestriction)
		out := out.(*GroupRestriction)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]meta_v1.LabelSelector, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*meta_v1.LabelSelector)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_IsPersonalSubjectAccessReview is an autogenerated deepcopy function.
func DeepCopy_v1_IsPersonalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IsPersonalSubjectAccessReview)
		out := out.(*IsPersonalSubjectAccessReview)
		*out = *in
		return nil
	}
}

// DeepCopy_v1_LocalResourceAccessReview is an autogenerated deepcopy function.
func DeepCopy_v1_LocalResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalResourceAccessReview)
		out := out.(*LocalResourceAccessReview)
		*out = *in
		if err := DeepCopy_v1_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_LocalSubjectAccessReview is an autogenerated deepcopy function.
func DeepCopy_v1_LocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		*out = *in
		if err := DeepCopy_v1_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		if in.GroupsSlice != nil {
			in, out := &in.GroupsSlice, &out.GroupsSlice
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make(OptionalScopes, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_NamedClusterRole is an autogenerated deepcopy function.
func DeepCopy_v1_NamedClusterRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamedClusterRole)
		out := out.(*NamedClusterRole)
		*out = *in
		if err := DeepCopy_v1_ClusterRole(&in.Role, &out.Role, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_NamedClusterRoleBinding is an autogenerated deepcopy function.
func DeepCopy_v1_NamedClusterRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamedClusterRoleBinding)
		out := out.(*NamedClusterRoleBinding)
		*out = *in
		if err := DeepCopy_v1_ClusterRoleBinding(&in.RoleBinding, &out.RoleBinding, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_NamedRole is an autogenerated deepcopy function.
func DeepCopy_v1_NamedRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamedRole)
		out := out.(*NamedRole)
		*out = *in
		if err := DeepCopy_v1_Role(&in.Role, &out.Role, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_NamedRoleBinding is an autogenerated deepcopy function.
func DeepCopy_v1_NamedRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamedRoleBinding)
		out := out.(*NamedRoleBinding)
		*out = *in
		if err := DeepCopy_v1_RoleBinding(&in.RoleBinding, &out.RoleBinding, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_Policy is an autogenerated deepcopy function.
func DeepCopy_v1_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(NamedRoles, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_NamedRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_PolicyBinding is an autogenerated deepcopy function.
func DeepCopy_v1_PolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBinding)
		out := out.(*PolicyBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(NamedRoleBindings, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_NamedRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_PolicyBindingList is an autogenerated deepcopy function.
func DeepCopy_v1_PolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBindingList)
		out := out.(*PolicyBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_PolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_PolicyList is an autogenerated deepcopy function.
func DeepCopy_v1_PolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyList)
		out := out.(*PolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Policy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Policy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_PolicyRule is an autogenerated deepcopy function.
func DeepCopy_v1_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if newVal, err := c.DeepCopy(&in.AttributeRestrictions); err != nil {
			return err
		} else {
			out.AttributeRestrictions = *newVal.(*runtime.RawExtension)
		}
		if in.APIGroups != nil {
			in, out := &in.APIGroups, &out.APIGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ResourceNames != nil {
			in, out := &in.ResourceNames, &out.ResourceNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.NonResourceURLsSlice != nil {
			in, out := &in.NonResourceURLsSlice, &out.NonResourceURLsSlice
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_ResourceAccessReview is an autogenerated deepcopy function.
func DeepCopy_v1_ResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReview)
		out := out.(*ResourceAccessReview)
		*out = *in
		if err := DeepCopy_v1_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_ResourceAccessReviewResponse is an autogenerated deepcopy function.
func DeepCopy_v1_ResourceAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReviewResponse)
		out := out.(*ResourceAccessReviewResponse)
		*out = *in
		if in.UsersSlice != nil {
			in, out := &in.UsersSlice, &out.UsersSlice
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.GroupsSlice != nil {
			in, out := &in.GroupsSlice, &out.GroupsSlice
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_Role is an autogenerated deepcopy function.
func DeepCopy_v1_Role(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Role)
		out := out.(*Role)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_RoleBinding is an autogenerated deepcopy function.
func DeepCopy_v1_RoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBinding)
		out := out.(*RoleBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if in.UserNames != nil {
			in, out := &in.UserNames, &out.UserNames
			*out = make(OptionalNames, len(*in))
			copy(*out, *in)
		}
		if in.GroupNames != nil {
			in, out := &in.GroupNames, &out.GroupNames
			*out = make(OptionalNames, len(*in))
			copy(*out, *in)
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]api_v1.ObjectReference, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_RoleBindingList is an autogenerated deepcopy function.
func DeepCopy_v1_RoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingList)
		out := out.(*RoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_RoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_RoleBindingRestriction is an autogenerated deepcopy function.
func DeepCopy_v1_RoleBindingRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestriction)
		out := out.(*RoleBindingRestriction)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1_RoleBindingRestrictionSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_RoleBindingRestrictionList is an autogenerated deepcopy function.
func DeepCopy_v1_RoleBindingRestrictionList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionList)
		out := out.(*RoleBindingRestrictionList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBindingRestriction, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_RoleBindingRestriction(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_RoleBindingRestrictionSpec is an autogenerated deepcopy function.
func DeepCopy_v1_RoleBindingRestrictionSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionSpec)
		out := out.(*RoleBindingRestrictionSpec)
		*out = *in
		if in.UserRestriction != nil {
			in, out := &in.UserRestriction, &out.UserRestriction
			*out = new(UserRestriction)
			if err := DeepCopy_v1_UserRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		if in.GroupRestriction != nil {
			in, out := &in.GroupRestriction, &out.GroupRestriction
			*out = new(GroupRestriction)
			if err := DeepCopy_v1_GroupRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ServiceAccountRestriction != nil {
			in, out := &in.ServiceAccountRestriction, &out.ServiceAccountRestriction
			*out = new(ServiceAccountRestriction)
			if err := DeepCopy_v1_ServiceAccountRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1_RoleList is an autogenerated deepcopy function.
func DeepCopy_v1_RoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleList)
		out := out.(*RoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Role, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Role(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_SelfSubjectRulesReview is an autogenerated deepcopy function.
func DeepCopy_v1_SelfSubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReview)
		out := out.(*SelfSubjectRulesReview)
		*out = *in
		if err := DeepCopy_v1_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_SubjectRulesReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_SelfSubjectRulesReviewSpec is an autogenerated deepcopy function.
func DeepCopy_v1_SelfSubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReviewSpec)
		out := out.(*SelfSubjectRulesReviewSpec)
		*out = *in
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make(OptionalScopes, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_ServiceAccountReference is an autogenerated deepcopy function.
func DeepCopy_v1_ServiceAccountReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountReference)
		out := out.(*ServiceAccountReference)
		*out = *in
		return nil
	}
}

// DeepCopy_v1_ServiceAccountRestriction is an autogenerated deepcopy function.
func DeepCopy_v1_ServiceAccountRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountRestriction)
		out := out.(*ServiceAccountRestriction)
		*out = *in
		if in.ServiceAccounts != nil {
			in, out := &in.ServiceAccounts, &out.ServiceAccounts
			*out = make([]ServiceAccountReference, len(*in))
			copy(*out, *in)
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_SubjectAccessReview is an autogenerated deepcopy function.
func DeepCopy_v1_SubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		*out = *in
		if err := DeepCopy_v1_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		if in.GroupsSlice != nil {
			in, out := &in.GroupsSlice, &out.GroupsSlice
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make(OptionalScopes, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_SubjectAccessReviewResponse is an autogenerated deepcopy function.
func DeepCopy_v1_SubjectAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewResponse)
		out := out.(*SubjectAccessReviewResponse)
		*out = *in
		return nil
	}
}

// DeepCopy_v1_SubjectRulesReview is an autogenerated deepcopy function.
func DeepCopy_v1_SubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReview)
		out := out.(*SubjectRulesReview)
		*out = *in
		if err := DeepCopy_v1_SubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_SubjectRulesReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_SubjectRulesReviewSpec is an autogenerated deepcopy function.
func DeepCopy_v1_SubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewSpec)
		out := out.(*SubjectRulesReviewSpec)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make(OptionalScopes, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1_SubjectRulesReviewStatus is an autogenerated deepcopy function.
func DeepCopy_v1_SubjectRulesReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewStatus)
		out := out.(*SubjectRulesReviewStatus)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_UserRestriction is an autogenerated deepcopy function.
func DeepCopy_v1_UserRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserRestriction)
		out := out.(*UserRestriction)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]meta_v1.LabelSelector, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*meta_v1.LabelSelector)
				}
			}
		}
		return nil
	}
}