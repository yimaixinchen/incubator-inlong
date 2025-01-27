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

package org.apache.inlong.manager.dao.mapper;

import org.apache.ibatis.annotations.Param;
import org.apache.inlong.manager.common.pojo.source.SourcePageRequest;
import org.apache.inlong.manager.dao.entity.StreamSourceEntity;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface StreamSourceEntityMapper {

    int insert(StreamSourceEntity record);

    int insertSelective(StreamSourceEntity record);

    StreamSourceEntity selectById(Integer id);

    StreamSourceEntity selectByIdForUpdate(Integer id);

    /**
     * According to the inlong group id and inlong stream id, query the number of valid source
     */
    int selectCount(@Param("groupId") String groupId, @Param("streamId") String streamId);

    /**
     * Paging query source list based on conditions
     */
    List<StreamSourceEntity> selectByCondition(@Param("request") SourcePageRequest request);

    /**
     * Query valid source list by the given group id, stream id and source name.
     */
    List<StreamSourceEntity> selectByRelatedId(@Param("groupId") String groupId, @Param("streamId") String streamId,
            @Param("sourceName") String sourceName);

    /**
     * According to the group id, stream id and source type, query valid source entity list.
     */
    List<StreamSourceEntity> selectByRelatedIdAndTypeForUpdate(@Param("groupId") String groupId,
            @Param("streamId") String streamId, @Param("sourceType") String sourceType);

    /**
     * Query the tasks that need to be added.
     */
    List<StreamSourceEntity> selectByStatusForUpdate(@Param("list") List<Integer> list);

    /**
     * Query the sources with status 20x by the given agent IP and agent UUID.
     */
    List<StreamSourceEntity> selectByStatusAndIpForUpdate(@Param("statusList") List<Integer> statusList,
            @Param("agentIp") String agentIp, @Param("uuid") String uuid);

    /**
     * Get the distinct source type from the given groupId and streamId
     */
    List<String> selectSourceType(@Param("groupId") String groupId, @Param("streamId") String streamId);

    /**
     * Get all sources in temporary status.
     *
     * @apiNote Do not need to filter sources that is_deleted > 0.
     */
    List<StreamSourceEntity> selectTempStatusSource();

    int updateByPrimaryKeySelective(StreamSourceEntity record);

    int updateByPrimaryKey(StreamSourceEntity record);

    /**
     * Update the status to `nextStatus` by the given id.
     *
     * @apiNote Should not change the modify_time
     */
    int updateStatus(@Param("id") Integer id, @Param("nextStatus") Integer nextStatus);

    /**
     * Update the status to `nextStatus` by the given group id and stream id.
     *
     * @apiNote Should not change the modify_time
     */
    int updateStatusByRelatedId(@Param("groupId") String groupId, @Param("streamId") String streamId,
            @Param("nextStatus") Integer nextStatus);

    /**
     * Update the agentIp and uuid.
     *
     * @apiNote Should not change the modify_time
     */
    int updateIpAndUuid(@Param("id") Integer id, @Param("agentIp") String agentIp, @Param("uuid") String uuid);

    int updateSnapshot(StreamSourceEntity entity);

    int deleteByRelatedId(@Param("groupId") String groupId, @Param("streamId") String streamId);

}