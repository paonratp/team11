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
/**
 * 
 * @export
 * @interface ControllersPreemption
 */
export interface ControllersPreemption {
    /**
     * 
     * @type {string}
     * @memberof ControllersPreemption
     */
    otherpeopleid?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPreemption
     */
    otherpeoplephone?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPreemption
     */
    phoneuser?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPreemption
     */
    purpose?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPreemption
     */
    roominfo?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPreemption
     */
    user?: number;
}

export function ControllersPreemptionFromJSON(json: any): ControllersPreemption {
    return ControllersPreemptionFromJSONTyped(json, false);
}

export function ControllersPreemptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPreemption {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'otherpeopleid': !exists(json, 'otherpeopleid') ? undefined : json['otherpeopleid'],
        'otherpeoplephone': !exists(json, 'otherpeoplephone') ? undefined : json['otherpeoplephone'],
        'phoneuser': !exists(json, 'phoneuser') ? undefined : json['phoneuser'],
        'purpose': !exists(json, 'purpose') ? undefined : json['purpose'],
        'roominfo': !exists(json, 'roominfo') ? undefined : json['roominfo'],
        'user': !exists(json, 'user') ? undefined : json['user'],
    };
}

export function ControllersPreemptionToJSON(value?: ControllersPreemption | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'otherpeopleid': value.otherpeopleid,
        'otherpeoplephone': value.otherpeoplephone,
        'phoneuser': value.phoneuser,
        'purpose': value.purpose,
        'roominfo': value.roominfo,
        'user': value.user,
    };
}


