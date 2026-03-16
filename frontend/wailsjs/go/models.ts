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
	export class Asset {
	    id: string;
	    originalPath: string;
	    checksum: string;
	    type: string;
	    createdAt: string;
	    updatedAt: string;
	    thumbUrl?: string;
	
	    static createFrom(source: any = {}) {
	        return new Asset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.originalPath = source["originalPath"];
	        this.checksum = source["checksum"];
	        this.type = source["type"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.thumbUrl = source["thumbUrl"];
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

