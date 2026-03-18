export namespace models {
	
	export class Album {
	    id: string;
	    albumName: string;
	    description: string;
	    assetCount: number;
	    albumThumbnailAssetId: string;
	
	    static createFrom(source: any = {}) {
	        return new Album(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.albumName = source["albumName"];
	        this.description = source["description"];
	        this.assetCount = source["assetCount"];
	        this.albumThumbnailAssetId = source["albumThumbnailAssetId"];
	    }
	}
	export class ExifInfo {
	    fileSizeInByte: number;
	    exifImageWidth: number;
	    exifImageHeight: number;
	    make: string;
	    model: string;
	    lensModel: string;
	    fNumber: number;
	    focalLength: number;
	    iso: number;
	    exposureTime: string;
	    latitude?: number;
	    longitude?: number;
	    city: string;
	    state: string;
	    country: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new ExifInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileSizeInByte = source["fileSizeInByte"];
	        this.exifImageWidth = source["exifImageWidth"];
	        this.exifImageHeight = source["exifImageHeight"];
	        this.make = source["make"];
	        this.model = source["model"];
	        this.lensModel = source["lensModel"];
	        this.fNumber = source["fNumber"];
	        this.focalLength = source["focalLength"];
	        this.iso = source["iso"];
	        this.exposureTime = source["exposureTime"];
	        this.latitude = source["latitude"];
	        this.longitude = source["longitude"];
	        this.city = source["city"];
	        this.state = source["state"];
	        this.country = source["country"];
	        this.description = source["description"];
	    }
	}
	export class Asset {
	    id: string;
	    originalPath: string;
	    originalFileName: string;
	    checksum: string;
	    type: string;
	    fileCreatedAt: string;
	    fileModifiedAt: string;
	    localDateTime: string;
	    duration: string;
	    isFavorite: boolean;
	    createdAt: string;
	    updatedAt: string;
	    thumbUrl?: string;
	    exifInfo?: ExifInfo;
	
	    static createFrom(source: any = {}) {
	        return new Asset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.originalPath = source["originalPath"];
	        this.originalFileName = source["originalFileName"];
	        this.checksum = source["checksum"];
	        this.type = source["type"];
	        this.fileCreatedAt = source["fileCreatedAt"];
	        this.fileModifiedAt = source["fileModifiedAt"];
	        this.localDateTime = source["localDateTime"];
	        this.duration = source["duration"];
	        this.isFavorite = source["isFavorite"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.thumbUrl = source["thumbUrl"];
	        this.exifInfo = this.convertValues(source["exifInfo"], ExifInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class UploadQueueItem {
	    id: number;
	    filePath: string;
	    status: string;
	    retryCount: number;
	    lastAttempt?: string;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new UploadQueueItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.filePath = source["filePath"];
	        this.status = source["status"];
	        this.retryCount = source["retryCount"];
	        this.lastAttempt = source["lastAttempt"];
	        this.error = source["error"];
	    }
	}
	export class User {
	    id: string;
	    email: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.email = source["email"];
	        this.name = source["name"];
	    }
	}

}

