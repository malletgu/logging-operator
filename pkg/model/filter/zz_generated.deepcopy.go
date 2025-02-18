// +build !ignore_autogenerated

// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package filter

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedotFilterConfig) DeepCopyInto(out *DedotFilterConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedotFilterConfig.
func (in *DedotFilterConfig) DeepCopy() *DedotFilterConfig {
	if in == nil {
		return nil
	}
	out := new(DedotFilterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoIP) DeepCopyInto(out *GeoIP) {
	*out = *in
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]Record, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(Record, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoIP.
func (in *GeoIP) DeepCopy() *GeoIP {
	if in == nil {
		return nil
	}
	out := new(GeoIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParseSection) DeepCopyInto(out *ParseSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParseSection.
func (in *ParseSection) DeepCopy() *ParseSection {
	if in == nil {
		return nil
	}
	out := new(ParseSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParserConfig) DeepCopyInto(out *ParserConfig) {
	*out = *in
	if in.Parsers != nil {
		in, out := &in.Parsers, &out.Parsers
		*out = make([]ParseSection, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParserConfig.
func (in *ParserConfig) DeepCopy() *ParserConfig {
	if in == nil {
		return nil
	}
	out := new(ParserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordTransformer) DeepCopyInto(out *RecordTransformer) {
	*out = *in
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]Record, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(Record, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordTransformer.
func (in *RecordTransformer) DeepCopy() *RecordTransformer {
	if in == nil {
		return nil
	}
	out := new(RecordTransformer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StdOutFilterConfig) DeepCopyInto(out *StdOutFilterConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StdOutFilterConfig.
func (in *StdOutFilterConfig) DeepCopy() *StdOutFilterConfig {
	if in == nil {
		return nil
	}
	out := new(StdOutFilterConfig)
	in.DeepCopyInto(out)
	return out
}
