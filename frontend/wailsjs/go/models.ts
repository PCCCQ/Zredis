export namespace data_hash {
	
	export class HashValue {
	    key?: string;
	    value?: string;
	
	    static createFrom(source: any = {}) {
	        return new HashValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class HashData {
	    key?: string;
	    valueList?: HashValue[];
	
	    static createFrom(source: any = {}) {
	        return new HashData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.valueList = this.convertValues(source["valueList"], HashValue);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

}

export namespace data_list {
	
	export class ListData {
	    key?: string;
	    value?: any;
	    ttl?: number;
	    index?: number;
	    addTab?: string;
	
	    static createFrom(source: any = {}) {
	        return new ListData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.ttl = source["ttl"];
	        this.index = source["index"];
	        this.addTab = source["addTab"];
	    }
	}

}

export namespace data_set {
	
	export class SetData {
	    key?: string;
	    value?: any;
	    ttl?: number;
	
	    static createFrom(source: any = {}) {
	        return new SetData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.ttl = source["ttl"];
	    }
	}

}

export namespace data_string {
	
	export class StringData {
	    key?: string;
	    value?: any;
	    ttl?: number;
	
	    static createFrom(source: any = {}) {
	        return new StringData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.ttl = source["ttl"];
	    }
	}

}

export namespace data_zset {
	
	export class ZsetData {
	    key?: string;
	    valueList?: redis.Z[];
	    // Go type: redis
	    value?: any;
	
	    static createFrom(source: any = {}) {
	        return new ZsetData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.valueList = this.convertValues(source["valueList"], redis.Z);
	        this.value = this.convertValues(source["value"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

}

export namespace model {
	
	export class SimpleResponse {
	    code: number;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new SimpleResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.data = source["data"];
	    }
	}

}

export namespace redis_init {
	
	export class RedisClient {
	    IP?: string;
	    port?: string;
	    username?: string;
	    password?: string;
	    DB?: number;
	
	    static createFrom(source: any = {}) {
	        return new RedisClient(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IP = source["IP"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.DB = source["DB"];
	    }
	}

}

