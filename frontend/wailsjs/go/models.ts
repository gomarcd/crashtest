export namespace main {
	
	export class APIResponse {
	    statusCode: number;
	    headers: Record<string, string>;
	    body: any;
	    timeMs: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new APIResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.statusCode = source["statusCode"];
	        this.headers = source["headers"];
	        this.body = source["body"];
	        this.timeMs = source["timeMs"];
	        this.error = source["error"];
	    }
	}
	export class RequestConfig {
	    method: string;
	    url: string;
	    headers: Record<string, string>;
	    queryParams: Record<string, string>;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.method = source["method"];
	        this.url = source["url"];
	        this.headers = source["headers"];
	        this.queryParams = source["queryParams"];
	        this.body = source["body"];
	    }
	}

}

