export namespace main {
	
	export class TestObj {
	
	
	    static createFrom(source: any = {}) {
	        return new TestObj(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

