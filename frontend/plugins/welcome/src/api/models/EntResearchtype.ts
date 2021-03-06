/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntResearchtypeEdges,
    EntResearchtypeEdgesFromJSON,
    EntResearchtypeEdgesFromJSONTyped,
    EntResearchtypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntResearchtype
 */
export interface EntResearchtype {
    /**
     * TYPENAME holds the value of the "TYPE_NAME" field.
     * @type {string}
     * @memberof EntResearchtype
     */
    tYPENAME?: string;
    /**
     * 
     * @type {EntResearchtypeEdges}
     * @memberof EntResearchtype
     */
    edges?: EntResearchtypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntResearchtype
     */
    id?: number;
}

export function EntResearchtypeFromJSON(json: any): EntResearchtype {
    return EntResearchtypeFromJSONTyped(json, false);
}

export function EntResearchtypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntResearchtype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'tYPENAME': !exists(json, 'TYPE_NAME') ? undefined : json['TYPE_NAME'],
        'edges': !exists(json, 'edges') ? undefined : EntResearchtypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntResearchtypeToJSON(value?: EntResearchtype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'TYPE_NAME': value.tYPENAME,
        'edges': EntResearchtypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


