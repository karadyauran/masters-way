// @ts-nocheck
/* eslint-disable */
/**
 * Masters way API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SchemasCreateProblemPayload
 */
export interface SchemasCreateProblemPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateProblemPayload
     */
    dayReportUuid?: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateProblemPayload
     */
    description?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasCreateProblemPayload
     */
    isDone?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateProblemPayload
     */
    ownerUuid?: string;
}

/**
 * Check if a given object implements the SchemasCreateProblemPayload interface.
 */
export function instanceOfSchemasCreateProblemPayload(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasCreateProblemPayloadFromJSON(json: any): SchemasCreateProblemPayload {
    return SchemasCreateProblemPayloadFromJSONTyped(json, false);
}

export function SchemasCreateProblemPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreateProblemPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dayReportUuid': !exists(json, 'dayReportUuid') ? undefined : json['dayReportUuid'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'isDone': !exists(json, 'isDone') ? undefined : json['isDone'],
        'ownerUuid': !exists(json, 'ownerUuid') ? undefined : json['ownerUuid'],
    };
}


export function SchemasCreateProblemPayloadToJSON(value?: SchemasCreateProblemPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dayReportUuid': value.dayReportUuid,
        'description': value.description,
        'isDone': value.isDone,
        'ownerUuid': value.ownerUuid,
    };
}

