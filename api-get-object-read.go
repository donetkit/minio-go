/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2015-2020 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package minio

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/s3utils"
	"io"
	"net/http"
)

//GetObjectRead wrapper function that accepts a request context
func (c *Client) GetObjectRead(ctx context.Context, bucketName, objectName string, opts GetObjectOptions) (io.ReadCloser, ObjectInfo, http.Header, error) {
	// Input validation.
	if err := s3utils.CheckValidBucketName(bucketName); err != nil {
		return nil, ObjectInfo{}, nil, err
	}
	if err := s3utils.CheckValidObjectName(objectName); err != nil {
		return nil, ObjectInfo{}, nil, err
	}

	return c.getObject(ctx, bucketName, objectName, opts)

}
