/* tslint:disable */
/* eslint-disable */
/**
 * 
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
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
 * @interface IncidentsCreateRequestBody
 */
export interface IncidentsCreateRequestBody {
    /**
     * Name of the incident
     * @type {string}
     * @memberof IncidentsCreateRequestBody
     */
    name: string;
}

export function IncidentsCreateRequestBodyFromJSON(json: any): IncidentsCreateRequestBody {
    return IncidentsCreateRequestBodyFromJSONTyped(json, false);
}

export function IncidentsCreateRequestBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): IncidentsCreateRequestBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
    };
}

export function IncidentsCreateRequestBodyToJSON(value?: IncidentsCreateRequestBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
    };
}

