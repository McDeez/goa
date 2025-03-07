package testdata

const PayloadWithNestedTypesServerTypeCode = `// NewMethodPayloadWithNestedTypesPayload builds the payload of the
// "MethodPayloadWithNestedTypes" endpoint of the
// "ServicePayloadWithNestedTypes" service from the gRPC request type.
func NewMethodPayloadWithNestedTypesPayload(message *service_payload_with_nested_typespb.MethodPayloadWithNestedTypesRequest) *servicepayloadwithnestedtypes.MethodPayloadWithNestedTypesPayload {
	v := &servicepayloadwithnestedtypes.MethodPayloadWithNestedTypesPayload{}
	if message.AParams != nil {
		v.AParams = protobufServicePayloadWithNestedTypespbAParamsToServicepayloadwithnestedtypesAParams(message.AParams)
	}
	if message.BParams != nil {
		v.BParams = protobufServicePayloadWithNestedTypespbBParamsToServicepayloadwithnestedtypesBParams(message.BParams)
	}
	return v
}

// NewProtoMethodPayloadWithNestedTypesResponse builds the gRPC response type
// from the result of the "MethodPayloadWithNestedTypes" endpoint of the
// "ServicePayloadWithNestedTypes" service.
func NewProtoMethodPayloadWithNestedTypesResponse() *service_payload_with_nested_typespb.MethodPayloadWithNestedTypesResponse {
	message := &service_payload_with_nested_typespb.MethodPayloadWithNestedTypesResponse{}
	return message
}

// protobufServicePayloadWithNestedTypespbAParamsToServicepayloadwithnestedtypesAParams
// builds a value of type *servicepayloadwithnestedtypes.AParams from a value
// of type *service_payload_with_nested_typespb.AParams.
func protobufServicePayloadWithNestedTypespbAParamsToServicepayloadwithnestedtypesAParams(v *service_payload_with_nested_typespb.AParams) *servicepayloadwithnestedtypes.AParams {
	if v == nil {
		return nil
	}
	res := &servicepayloadwithnestedtypes.AParams{}
	if v.A != nil {
		res.A = make(map[string][]string, len(v.A))
		for key, val := range v.A {
			tk := key
			tv := make([]string, len(val.Field))
			for i, val := range val.Field {
				tv[i] = val
			}
			res.A[tk] = tv
		}
	}

	return res
}

// protobufServicePayloadWithNestedTypespbBParamsToServicepayloadwithnestedtypesBParams
// builds a value of type *servicepayloadwithnestedtypes.BParams from a value
// of type *service_payload_with_nested_typespb.BParams.
func protobufServicePayloadWithNestedTypespbBParamsToServicepayloadwithnestedtypesBParams(v *service_payload_with_nested_typespb.BParams) *servicepayloadwithnestedtypes.BParams {
	if v == nil {
		return nil
	}
	res := &servicepayloadwithnestedtypes.BParams{}
	if v.B != nil {
		res.B = make(map[string]string, len(v.B))
		for key, val := range v.B {
			tk := key
			tv := val
			res.B[tk] = tv
		}
	}

	return res
}

// svcServicepayloadwithnestedtypesAParamsToServicePayloadWithNestedTypespbAParams
// builds a value of type *service_payload_with_nested_typespb.AParams from a
// value of type *servicepayloadwithnestedtypes.AParams.
func svcServicepayloadwithnestedtypesAParamsToServicePayloadWithNestedTypespbAParams(v *servicepayloadwithnestedtypes.AParams) *service_payload_with_nested_typespb.AParams {
	if v == nil {
		return nil
	}
	res := &service_payload_with_nested_typespb.AParams{}
	if v.A != nil {
		res.A = make(map[string]*service_payload_with_nested_typespb.ArrayOfString, len(v.A))
		for key, val := range v.A {
			tk := key
			tv := &service_payload_with_nested_typespb.ArrayOfString{}
			tv.Field = make([]string, len(val))
			for i, val := range val {
				tv.Field[i] = val
			}
			res.A[tk] = tv
		}
	}

	return res
}

// svcServicepayloadwithnestedtypesBParamsToServicePayloadWithNestedTypespbBParams
// builds a value of type *service_payload_with_nested_typespb.BParams from a
// value of type *servicepayloadwithnestedtypes.BParams.
func svcServicepayloadwithnestedtypesBParamsToServicePayloadWithNestedTypespbBParams(v *servicepayloadwithnestedtypes.BParams) *service_payload_with_nested_typespb.BParams {
	if v == nil {
		return nil
	}
	res := &service_payload_with_nested_typespb.BParams{}
	if v.B != nil {
		res.B = make(map[string]string, len(v.B))
		for key, val := range v.B {
			tk := key
			tv := val
			res.B[tk] = tv
		}
	}

	return res
}
`

const PayloadWithMultipleUseTypesServerTypeCode = `// NewMethodPayloadDuplicateAPayload builds the payload of the
// "MethodPayloadDuplicateA" endpoint of the "ServicePayloadWithNestedTypes"
// service from the gRPC request type.
func NewMethodPayloadDuplicateAPayload(message *service_payload_with_nested_typespb.DupePayload) servicepayloadwithnestedtypes.DupePayload {
	v := servicepayloadwithnestedtypes.DupePayload(message.Field)
	return v
}

// NewProtoMethodPayloadDuplicateAResponse builds the gRPC response type from
// the result of the "MethodPayloadDuplicateA" endpoint of the
// "ServicePayloadWithNestedTypes" service.
func NewProtoMethodPayloadDuplicateAResponse() *service_payload_with_nested_typespb.MethodPayloadDuplicateAResponse {
	message := &service_payload_with_nested_typespb.MethodPayloadDuplicateAResponse{}
	return message
}

// NewMethodPayloadDuplicateBPayload builds the payload of the
// "MethodPayloadDuplicateB" endpoint of the "ServicePayloadWithNestedTypes"
// service from the gRPC request type.
func NewMethodPayloadDuplicateBPayload(message *service_payload_with_nested_typespb.DupePayload) servicepayloadwithnestedtypes.DupePayload {
	v := servicepayloadwithnestedtypes.DupePayload(message.Field)
	return v
}

// NewProtoMethodPayloadDuplicateBResponse builds the gRPC response type from
// the result of the "MethodPayloadDuplicateB" endpoint of the
// "ServicePayloadWithNestedTypes" service.
func NewProtoMethodPayloadDuplicateBResponse() *service_payload_with_nested_typespb.MethodPayloadDuplicateBResponse {
	message := &service_payload_with_nested_typespb.MethodPayloadDuplicateBResponse{}
	return message
}
`

const PayloadWithAliasTypeServerTypeCode = `// NewMethodMessageUserTypeWithAliasPayload builds the payload of the
// "MethodMessageUserTypeWithAlias" endpoint of the
// "ServiceMessageUserTypeWithAlias" service from the gRPC request type.
func NewMethodMessageUserTypeWithAliasPayload(message *service_message_user_type_with_aliaspb.MethodMessageUserTypeWithAliasRequest) *servicemessageusertypewithalias.PayloadAliasT {
	v := &servicemessageusertypewithalias.PayloadAliasT{
		IntAliasField: servicemessageusertypewithalias.IntAlias(message.IntAliasField),
	}
	if message.OptionalIntAliasField != 0 {
		optionalIntAliasFieldptr := servicemessageusertypewithalias.IntAlias(message.OptionalIntAliasField)
		v.OptionalIntAliasField = &optionalIntAliasFieldptr
	}
	return v
}

// NewProtoMethodMessageUserTypeWithAliasResponse builds the gRPC response type
// from the result of the "MethodMessageUserTypeWithAlias" endpoint of the
// "ServiceMessageUserTypeWithAlias" service.
func NewProtoMethodMessageUserTypeWithAliasResponse(result *servicemessageusertypewithalias.PayloadAliasT) *service_message_user_type_with_aliaspb.MethodMessageUserTypeWithAliasResponse {
	message := &service_message_user_type_with_aliaspb.MethodMessageUserTypeWithAliasResponse{
		IntAliasField: int32(result.IntAliasField),
	}
	if result.OptionalIntAliasField != nil {
		message.OptionalIntAliasField = int32(*result.OptionalIntAliasField)
	}
	return message
}
`

const ResultWithCollectionServerTypeCode = `// NewProtoMethodResultWithCollectionResponse builds the gRPC response type
// from the result of the "MethodResultWithCollection" endpoint of the
// "ServiceResultWithCollection" service.
func NewProtoMethodResultWithCollectionResponse(result *serviceresultwithcollection.MethodResultWithCollectionResult) *service_result_with_collectionpb.MethodResultWithCollectionResponse {
	message := &service_result_with_collectionpb.MethodResultWithCollectionResponse{}
	if result.Result != nil {
		message.Result = svcServiceresultwithcollectionResultTToServiceResultWithCollectionpbResultT(result.Result)
	}
	return message
}

// svcServiceresultwithcollectionResultTToServiceResultWithCollectionpbResultT
// builds a value of type *service_result_with_collectionpb.ResultT from a
// value of type *serviceresultwithcollection.ResultT.
func svcServiceresultwithcollectionResultTToServiceResultWithCollectionpbResultT(v *serviceresultwithcollection.ResultT) *service_result_with_collectionpb.ResultT {
	if v == nil {
		return nil
	}
	res := &service_result_with_collectionpb.ResultT{}
	if v.CollectionField != nil {
		res.CollectionField = &service_result_with_collectionpb.RTCollection{}
		res.CollectionField.Field = make([]*service_result_with_collectionpb.RT, len(v.CollectionField))
		for i, val := range v.CollectionField {
			res.CollectionField.Field[i] = &service_result_with_collectionpb.RT{}
			if val.IntField != nil {
				res.CollectionField.Field[i].IntField = int32(*val.IntField)
			}
		}
	}

	return res
}

// protobufServiceResultWithCollectionpbResultTToServiceresultwithcollectionResultT
// builds a value of type *serviceresultwithcollection.ResultT from a value of
// type *service_result_with_collectionpb.ResultT.
func protobufServiceResultWithCollectionpbResultTToServiceresultwithcollectionResultT(v *service_result_with_collectionpb.ResultT) *serviceresultwithcollection.ResultT {
	if v == nil {
		return nil
	}
	res := &serviceresultwithcollection.ResultT{}
	if v.CollectionField != nil {
		res.CollectionField = make([]*serviceresultwithcollection.RT, len(v.CollectionField.Field))
		for i, val := range v.CollectionField.Field {
			res.CollectionField[i] = &serviceresultwithcollection.RT{}
			if val.IntField != 0 {
				intFieldptr := int(val.IntField)
				res.CollectionField[i].IntField = &intFieldptr
			}
		}
	}

	return res
}
`

const PayloadWithMixedAttributesServerTypeCode = `// NewUnaryMethodPayload builds the payload of the "UnaryMethod" endpoint of
// the "ServicePayloadWithMixedAttributes" service from the gRPC request type.
func NewUnaryMethodPayload(message *service_payload_with_mixed_attributespb.UnaryMethodRequest) *servicepayloadwithmixedattributes.APayload {
	v := &servicepayloadwithmixedattributes.APayload{
		Required:        int(message.Required),
		Default:         int(message.Default),
		RequiredDefault: int(message.RequiredDefault),
	}
	if message.Optional != 0 {
		optionalptr := int(message.Optional)
		v.Optional = &optionalptr
	}
	if message.Default == 0 {
		v.Default = 100
	}
	return v
}

// NewProtoUnaryMethodResponse builds the gRPC response type from the result of
// the "UnaryMethod" endpoint of the "ServicePayloadWithMixedAttributes"
// service.
func NewProtoUnaryMethodResponse() *service_payload_with_mixed_attributespb.UnaryMethodResponse {
	message := &service_payload_with_mixed_attributespb.UnaryMethodResponse{}
	return message
}

// NewProtoStreamingMethodResponse builds the gRPC response type from the
// result of the "StreamingMethod" endpoint of the
// "ServicePayloadWithMixedAttributes" service.
func NewProtoStreamingMethodResponse() *service_payload_with_mixed_attributespb.StreamingMethodResponse {
	message := &service_payload_with_mixed_attributespb.StreamingMethodResponse{}
	return message
}

func NewAPayload(v *service_payload_with_mixed_attributespb.StreamingMethodStreamingRequest) *servicepayloadwithmixedattributes.APayload {
	spayload := &servicepayloadwithmixedattributes.APayload{
		Required:        int(v.Required),
		Default:         int(v.Default),
		RequiredDefault: int(v.RequiredDefault),
	}
	if v.Optional != 0 {
		optionalptr := int(v.Optional)
		spayload.Optional = &optionalptr
	}
	if v.Default == 0 {
		spayload.Default = 100
	}
	return spayload
}
`

const WithErrorsServerTypeCode = `// NewMethodUnaryRPCWithErrorsPayload builds the payload of the
// "MethodUnaryRPCWithErrors" endpoint of the "ServiceUnaryRPCWithErrors"
// service from the gRPC request type.
func NewMethodUnaryRPCWithErrorsPayload(message *service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsRequest) string {
	v := message.Field
	return v
}

// NewProtoMethodUnaryRPCWithErrorsResponse builds the gRPC response type from
// the result of the "MethodUnaryRPCWithErrors" endpoint of the
// "ServiceUnaryRPCWithErrors" service.
func NewProtoMethodUnaryRPCWithErrorsResponse(result string) *service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsResponse {
	message := &service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsResponse{}
	message.Field = result
	return message
}

// NewMethodUnaryRPCWithErrorsInternalError builds the gRPC error response type
// from the error of the "MethodUnaryRPCWithErrors" endpoint of the
// "ServiceUnaryRPCWithErrors" service.
func NewMethodUnaryRPCWithErrorsInternalError(er *serviceunaryrpcwitherrors.AnotherError) *service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsInternalError {
	message := &service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsInternalError{
		Name: er.Name,
	}
	if er.Description != nil {
		message.Description = *er.Description
	}
	return message
}

// NewMethodUnaryRPCWithErrorsBadRequestError builds the gRPC error response
// type from the error of the "MethodUnaryRPCWithErrors" endpoint of the
// "ServiceUnaryRPCWithErrors" service.
func NewMethodUnaryRPCWithErrorsBadRequestError(er *serviceunaryrpcwitherrors.AnotherError) *service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsBadRequestError {
	message := &service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsBadRequestError{
		Name: er.Name,
	}
	if er.Description != nil {
		message.Description = *er.Description
	}
	return message
}

// NewMethodUnaryRPCWithErrorsCustomErrorError builds the gRPC error response
// type from the error of the "MethodUnaryRPCWithErrors" endpoint of the
// "ServiceUnaryRPCWithErrors" service.
func NewMethodUnaryRPCWithErrorsCustomErrorError(er *serviceunaryrpcwitherrors.ErrorType) *service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsCustomErrorError {
	message := &service_unary_rpc_with_errorspb.MethodUnaryRPCWithErrorsCustomErrorError{}
	if er.A != nil {
		message.A = *er.A
	}
	return message
}
`

const ElemValidationServerTypesFile = `// NewMethodElemValidationPayload builds the payload of the
// "MethodElemValidation" endpoint of the "ServiceElemValidation" service from
// the gRPC request type.
func NewMethodElemValidationPayload(message *service_elem_validationpb.MethodElemValidationRequest) *serviceelemvalidation.ResultType {
	v := &serviceelemvalidation.ResultType{}
	if message.Foo != nil {
		v.Foo = make(map[string][]string, len(message.Foo))
		for key, val := range message.Foo {
			tk := key
			tv := make([]string, len(val.Field))
			for i, val := range val.Field {
				tv[i] = val
			}
			v.Foo[tk] = tv
		}
	}
	return v
}

// NewProtoMethodElemValidationResponse builds the gRPC response type from the
// result of the "MethodElemValidation" endpoint of the "ServiceElemValidation"
// service.
func NewProtoMethodElemValidationResponse() *service_elem_validationpb.MethodElemValidationResponse {
	message := &service_elem_validationpb.MethodElemValidationResponse{}
	return message
}

// ValidateMethodElemValidationRequest runs the validations defined on
// MethodElemValidationRequest.
func ValidateMethodElemValidationRequest(message *service_elem_validationpb.MethodElemValidationRequest) (err error) {
	for _, v := range message.Foo {
		if v != nil {
			if err2 := ValidateArrayOfString(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateArrayOfString runs the validations defined on ArrayOfString.
func ValidateArrayOfString(val *service_elem_validationpb.ArrayOfString) (err error) {
	if len(val.Field) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("val.field", val.Field, len(val.Field), 1, true))
	}
	return
}
`

const AliasValidationServerTypesFile = `// NewMethodElemValidationPayload builds the payload of the
// "MethodElemValidation" endpoint of the "ServiceElemValidation" service from
// the gRPC request type.
func NewMethodElemValidationPayload(message *service_elem_validationpb.UUID) serviceelemvalidation.UUID {
	v := serviceelemvalidation.UUID(message.Field)
	return v
}

// NewProtoMethodElemValidationResponse builds the gRPC response type from the
// result of the "MethodElemValidation" endpoint of the "ServiceElemValidation"
// service.
func NewProtoMethodElemValidationResponse() *service_elem_validationpb.MethodElemValidationResponse {
	message := &service_elem_validationpb.MethodElemValidationResponse{}
	return message
}

// ValidateUUID runs the validations defined on UUID.
func ValidateUUID(message *service_elem_validationpb.UUID) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.field", message.Field, goa.FormatUUID))

	return
}
`

const StructMetaTypeServerTypeCode = `// NewMethodPayload builds the payload of the "Method" endpoint of the
// "UsingMetaTypes" service from the gRPC request type.
func NewMethodPayload(message *using_meta_typespb.MethodRequest) *usingmetatypes.MethodPayload {
	v := &usingmetatypes.MethodPayload{
		A: flag.ErrorHandling(message.A),
		B: flag.ErrorHandling(message.B),
	}
	if message.D != 0 {
		dptr := flag.ErrorHandling(message.D)
		v.D = &dptr
	}
	var zeroMessageA int64
	if message.A == zeroMessageA {
		v.A = 1
	}
	var zeroMessageB int64
	if message.B == zeroMessageB {
		v.B = 2
	}
	if message.C != nil {
		v.C = make([]time.Duration, len(message.C))
		for i, val := range message.C {
			v.C[i] = time.Duration(val)
		}
	}
	return v
}

// NewProtoMethodResponse builds the gRPC response type from the result of the
// "Method" endpoint of the "UsingMetaTypes" service.
func NewProtoMethodResponse(result *usingmetatypes.MethodResult) *using_meta_typespb.MethodResponse {
	message := &using_meta_typespb.MethodResponse{
		A: int64(result.A),
		B: int64(result.B),
	}
	if result.D != nil {
		message.D = int64(*result.D)
	}
	if result.C != nil {
		message.C = make([]int64, len(result.C))
		for i, val := range result.C {
			message.C[i] = int64(val)
		}
	}
	return message
}
`

const DefaultFieldsServerTypeCode = `// NewMethodPayload builds the payload of the "Method" endpoint of the
// "DefaultFields" service from the gRPC request type.
func NewMethodPayload(message *default_fieldspb.MethodRequest) *defaultfields.MethodPayload {
	v := &defaultfields.MethodPayload{
		Req:  message.Req,
		Def0: message.Def0,
		Def1: message.Def1,
		Def2: message.Def2,
		Reqs: message.Reqs,
		Defs: message.Defs,
		Defe: message.Defe,
		Rat:  message.Rat,
		Flt0: message.Flt0,
		Flt1: message.Flt1,
	}
	if message.Opt != 0 {
		v.Opt = &message.Opt
	}
	if message.Opts != "" {
		v.Opts = &message.Opts
	}
	if message.Flt != 0 {
		v.Flt = &message.Flt
	}
	if message.Def1 == 0 {
		v.Def1 = 1
	}
	if message.Def2 == 0 {
		v.Def2 = 2
	}
	if message.Defs == "" {
		v.Defs = "!"
	}
	if message.Flt1 == 0 {
		v.Flt1 = 1
	}
	return v
}

// NewProtoMethodResponse builds the gRPC response type from the result of the
// "Method" endpoint of the "DefaultFields" service.
func NewProtoMethodResponse() *default_fieldspb.MethodResponse {
	message := &default_fieldspb.MethodResponse{}
	return message
}
`
