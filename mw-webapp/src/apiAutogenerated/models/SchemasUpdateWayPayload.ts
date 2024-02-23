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
 * @interface SchemasUpdateWayPayload
 */
export interface SchemasUpdateWayPayload {
    /**
     * TODO estimationTime's type int 32?
     * @type {number}
     * @memberof SchemasUpdateWayPayload
     */
    estimationTime?: number;
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateWayPayload
     */
    goalDescription?: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateWayPayload
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateWayPayload
     */
    status?: string;
}

/**
 * Check if a given object implements the SchemasUpdateWayPayload interface.
 */
export function instanceOfSchemasUpdateWayPayload(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasUpdateWayPayloadFromJSON(json: any): SchemasUpdateWayPayload {
    return SchemasUpdateWayPayloadFromJSONTyped(json, false);
}

export function SchemasUpdateWayPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUpdateWayPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'estimationTime': !exists(json, 'estimationTime') ? undefined : json['estimationTime'],
        'goalDescription': !exists(json, 'goalDescription') ? undefined : json['goalDescription'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}


export function SchemasUpdateWayPayloadToJSON(value?: SchemasUpdateWayPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'estimationTime': value.estimationTime,
        'goalDescription': value.goalDescription,
        'name': value.name,
        'status': value.status,
    };
}

