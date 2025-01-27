/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.inlong.manager.client.api;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.apache.inlong.manager.common.enums.FieldType;

@Data
@NoArgsConstructor
@ApiModel("Sink field configuration")
public class SinkField extends StreamField {

    @ApiModelProperty("Source field name")
    private String sourceFieldName;

    @ApiModelProperty("Source field type")
    private String sourceFieldType;

    @ApiModelProperty("Is source meta field, 0: no, 1: yes")
    private Integer isSourceMetaField = 0;

    public SinkField(int index, FieldType fieldType, String fieldName, String fieldComment,
            String fieldValue, String sourceFieldName, String sourceFieldType, Integer isSourceMetaField) {
        super(index, fieldType, fieldName, fieldComment, fieldValue);
        this.sourceFieldName = sourceFieldName;
        this.sourceFieldType = sourceFieldType;
        this.isSourceMetaField = isSourceMetaField;
    }
}
