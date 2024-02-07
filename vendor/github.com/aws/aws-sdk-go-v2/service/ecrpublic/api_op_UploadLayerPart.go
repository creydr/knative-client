// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads an image layer part to Amazon ECR. When an image is pushed, each new
// image layer is uploaded in parts. The maximum size of each image layer part can
// be 20971520 bytes (about 20MB). The UploadLayerPart API is called once for each
// new image layer part. This operation is used by the Amazon ECR proxy and is not
// generally used by customers for pulling and pushing images. In most cases, you
// should use the docker CLI to pull, tag, and push images.
func (c *Client) UploadLayerPart(ctx context.Context, params *UploadLayerPartInput, optFns ...func(*Options)) (*UploadLayerPartOutput, error) {
	if params == nil {
		params = &UploadLayerPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadLayerPart", params, optFns, c.addOperationUploadLayerPartMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadLayerPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadLayerPartInput struct {

	// The base64-encoded layer part payload.
	//
	// This member is required.
	LayerPartBlob []byte

	// The position of the first byte of the layer part witin the overall image layer.
	//
	// This member is required.
	PartFirstByte *int64

	// The position of the last byte of the layer part within the overall image layer.
	//
	// This member is required.
	PartLastByte *int64

	// The name of the repository that you're uploading layer parts to.
	//
	// This member is required.
	RepositoryName *string

	// The upload ID from a previous InitiateLayerUpload operation to associate with
	// the layer part upload.
	//
	// This member is required.
	UploadId *string

	// The Amazon Web Services account ID, or registry alias, that's associated with
	// the registry that you're uploading layer parts to. If you do not specify a
	// registry, the default public registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type UploadLayerPartOutput struct {

	// The integer value of the last byte that's received in the request.
	LastByteReceived *int64

	// The registry ID that's associated with the request.
	RegistryId *string

	// The repository name that's associated with the request.
	RepositoryName *string

	// The upload ID that's associated with the request.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUploadLayerPartMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUploadLayerPart{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUploadLayerPart{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUploadLayerPartValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadLayerPart(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUploadLayerPart(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "UploadLayerPart",
	}
}
