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
 * @interface ControllersBook
 */
export interface ControllersBook {
    /**
     * 
     * @type {number}
     * @memberof ControllersBook
     */
    author?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersBook
     */
    barCode?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersBook
     */
    bookPage?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersBook
     */
    bookname?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersBook
     */
    category?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBook
     */
    userid?: number;
}

export function ControllersBookFromJSON(json: any): ControllersBook {
    return ControllersBookFromJSONTyped(json, false);
}

export function ControllersBookFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBook {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'author': !exists(json, 'author') ? undefined : json['author'],
        'barCode': !exists(json, 'barCode') ? undefined : json['barCode'],
        'bookPage': !exists(json, 'bookPage') ? undefined : json['bookPage'],
        'bookname': !exists(json, 'bookname') ? undefined : json['bookname'],
        'category': !exists(json, 'category') ? undefined : json['category'],
        'userid': !exists(json, 'userid') ? undefined : json['userid'],
    };
}

export function ControllersBookToJSON(value?: ControllersBook | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'author': value.author,
        'barCode': value.barCode,
        'bookPage': value.bookPage,
        'bookname': value.bookname,
        'category': value.category,
        'userid': value.userid,
    };
}


