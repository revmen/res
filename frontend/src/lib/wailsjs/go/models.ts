export namespace models {
	
	export class Candidate {
	    name: string;
	    x: number;
	    y: number;
	    power: number;
	
	    static createFrom(source: any = {}) {
	        return new Candidate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.power = source["power"];
	    }
	}
	export class CandidateGen {
	    name: string;
	    x: number;
	    y: number;
	    radius: number;
	    distribution: string;
	    power: number;
	
	    static createFrom(source: any = {}) {
	        return new CandidateGen(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.radius = source["radius"];
	        this.distribution = source["distribution"];
	        this.power = source["power"];
	    }
	}
	export class Voter {
	
	
	    static createFrom(source: any = {}) {
	        return new Voter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Frame {
	    candidates: Candidate[];
	    voters: Voter[];
	
	    static createFrom(source: any = {}) {
	        return new Frame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.candidates = this.convertValues(source["candidates"], Candidate);
	        this.voters = this.convertValues(source["voters"], Voter);
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
	export class VoterGen {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new VoterGen(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class Experiment {
	    name: string;
	    notes: string;
	    voterGens: VoterGen[];
	    candidateGens: CandidateGen[];
	    frames: Frame[];
	
	    static createFrom(source: any = {}) {
	        return new Experiment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.notes = source["notes"];
	        this.voterGens = this.convertValues(source["voterGens"], VoterGen);
	        this.candidateGens = this.convertValues(source["candidateGens"], CandidateGen);
	        this.frames = this.convertValues(source["frames"], Frame);
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
	

}

