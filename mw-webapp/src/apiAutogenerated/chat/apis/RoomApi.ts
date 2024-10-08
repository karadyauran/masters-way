// @ts-nocheck
/* eslint-disable */
/**
 * Masters way chat API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  SchemasCreateRoomPayload,
  SchemasGetRoomPreviewResponse,
  SchemasGetRoomsResponse,
  SchemasRoomPopulatedResponse,
} from '../models/index';
import {
    SchemasCreateRoomPayloadFromJSON,
    SchemasCreateRoomPayloadToJSON,
    SchemasGetRoomPreviewResponseFromJSON,
    SchemasGetRoomPreviewResponseToJSON,
    SchemasGetRoomsResponseFromJSON,
    SchemasGetRoomsResponseToJSON,
    SchemasRoomPopulatedResponseFromJSON,
    SchemasRoomPopulatedResponseToJSON,
} from '../models/index';

export interface AddUserToRoomRequest {
    roomId: string;
    userId: string;
}

export interface CreateRoomRequest {
    request: SchemasCreateRoomPayload;
}

export interface DeleteUserFromRoomRequest {
    roomId: string;
    userId: string;
}

export interface GetRoomByIdRequest {
    roomId: string;
}

export interface GetRoomsRequest {
    roomType: GetRoomsRoomTypeEnum;
}

export interface UpdateRoomRequest {
    roomId: string;
}

/**
 * 
 */
export class RoomApi extends runtime.BaseAPI {

    /**
     * Add user to room
     */
    async addUserToRoomRaw(requestParameters: AddUserToRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasRoomPopulatedResponse>> {
        if (requestParameters.roomId === null || requestParameters.roomId === undefined) {
            throw new runtime.RequiredError('roomId','Required parameter requestParameters.roomId was null or undefined when calling addUserToRoom.');
        }

        if (requestParameters.userId === null || requestParameters.userId === undefined) {
            throw new runtime.RequiredError('userId','Required parameter requestParameters.userId was null or undefined when calling addUserToRoom.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/{roomId}/users/{userId}`.replace(`{${"roomId"}}`, encodeURIComponent(String(requestParameters.roomId))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters.userId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasRoomPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Add user to room
     */
    async addUserToRoom(requestParameters: AddUserToRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasRoomPopulatedResponse> {
        const response = await this.addUserToRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create room for user
     */
    async createRoomRaw(requestParameters: CreateRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasRoomPopulatedResponse>> {
        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling createRoom.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/rooms`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasCreateRoomPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasRoomPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Create room for user
     */
    async createRoom(requestParameters: CreateRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasRoomPopulatedResponse> {
        const response = await this.createRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete user from room
     */
    async deleteUserFromRoomRaw(requestParameters: DeleteUserFromRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasRoomPopulatedResponse>> {
        if (requestParameters.roomId === null || requestParameters.roomId === undefined) {
            throw new runtime.RequiredError('roomId','Required parameter requestParameters.roomId was null or undefined when calling deleteUserFromRoom.');
        }

        if (requestParameters.userId === null || requestParameters.userId === undefined) {
            throw new runtime.RequiredError('userId','Required parameter requestParameters.userId was null or undefined when calling deleteUserFromRoom.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/{roomId}/users/{userId}`.replace(`{${"roomId"}}`, encodeURIComponent(String(requestParameters.roomId))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters.userId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasRoomPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Delete user from room
     */
    async deleteUserFromRoom(requestParameters: DeleteUserFromRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasRoomPopulatedResponse> {
        const response = await this.deleteUserFromRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get chat preview
     */
    async getChatPreviewRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasGetRoomPreviewResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/preview`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasGetRoomPreviewResponseFromJSON(jsonValue));
    }

    /**
     * Get chat preview
     */
    async getChatPreview(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasGetRoomPreviewResponse> {
        const response = await this.getChatPreviewRaw(initOverrides);
        return await response.value();
    }

    /**
     * Get room by id
     */
    async getRoomByIdRaw(requestParameters: GetRoomByIdRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasRoomPopulatedResponse>> {
        if (requestParameters.roomId === null || requestParameters.roomId === undefined) {
            throw new runtime.RequiredError('roomId','Required parameter requestParameters.roomId was null or undefined when calling getRoomById.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/{roomId}`.replace(`{${"roomId"}}`, encodeURIComponent(String(requestParameters.roomId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasRoomPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Get room by id
     */
    async getRoomById(requestParameters: GetRoomByIdRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasRoomPopulatedResponse> {
        const response = await this.getRoomByIdRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get rooms for user
     */
    async getRoomsRaw(requestParameters: GetRoomsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasGetRoomsResponse>> {
        if (requestParameters.roomType === null || requestParameters.roomType === undefined) {
            throw new runtime.RequiredError('roomType','Required parameter requestParameters.roomType was null or undefined when calling getRooms.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/list/{roomType}`.replace(`{${"roomType"}}`, encodeURIComponent(String(requestParameters.roomType))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasGetRoomsResponseFromJSON(jsonValue));
    }

    /**
     * Get rooms for user
     */
    async getRooms(requestParameters: GetRoomsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasGetRoomsResponse> {
        const response = await this.getRoomsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update room
     */
    async updateRoomRaw(requestParameters: UpdateRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasRoomPopulatedResponse>> {
        if (requestParameters.roomId === null || requestParameters.roomId === undefined) {
            throw new runtime.RequiredError('roomId','Required parameter requestParameters.roomId was null or undefined when calling updateRoom.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/{roomId}`.replace(`{${"roomId"}}`, encodeURIComponent(String(requestParameters.roomId))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasRoomPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Update room
     */
    async updateRoom(requestParameters: UpdateRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasRoomPopulatedResponse> {
        const response = await this.updateRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const GetRoomsRoomTypeEnum = {
    Private: 'private',
    Group: 'group'
} as const;
export type GetRoomsRoomTypeEnum = typeof GetRoomsRoomTypeEnum[keyof typeof GetRoomsRoomTypeEnum];
