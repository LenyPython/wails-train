export namespace testpkg {
	
	export class TestObj {
	    name: string;
	    job: string;
	
	    static createFrom(source: any = {}) {
	        return new TestObj(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.job = source["job"];
	    }
	}

}

