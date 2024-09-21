import * as jspb from 'google-protobuf'



export class CreateJobPostingRequest extends jspb.Message {
  getIndustry(): string;
  setIndustry(value: string): CreateJobPostingRequest;

  getCompanyName(): string;
  setCompanyName(value: string): CreateJobPostingRequest;

  getLocation(): string;
  setLocation(value: string): CreateJobPostingRequest;

  getTitle(): string;
  setTitle(value: string): CreateJobPostingRequest;

  getBaseSalary(): Salary | undefined;
  setBaseSalary(value?: Salary): CreateJobPostingRequest;
  hasBaseSalary(): boolean;
  clearBaseSalary(): CreateJobPostingRequest;

  getBonusSalary(): Salary | undefined;
  setBonusSalary(value?: Salary): CreateJobPostingRequest;
  hasBonusSalary(): boolean;
  clearBonusSalary(): CreateJobPostingRequest;

  getQualificationsList(): Array<string>;
  setQualificationsList(value: Array<string>): CreateJobPostingRequest;
  clearQualificationsList(): CreateJobPostingRequest;
  addQualifications(value: string, index?: number): CreateJobPostingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateJobPostingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateJobPostingRequest): CreateJobPostingRequest.AsObject;
  static serializeBinaryToWriter(message: CreateJobPostingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateJobPostingRequest;
  static deserializeBinaryFromReader(message: CreateJobPostingRequest, reader: jspb.BinaryReader): CreateJobPostingRequest;
}

export namespace CreateJobPostingRequest {
  export type AsObject = {
    industry: string,
    companyName: string,
    location: string,
    title: string,
    baseSalary?: Salary.AsObject,
    bonusSalary?: Salary.AsObject,
    qualificationsList: Array<string>,
  }
}

export class CreateJobPostingResponse extends jspb.Message {
  getJobPostingId(): string;
  setJobPostingId(value: string): CreateJobPostingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateJobPostingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateJobPostingResponse): CreateJobPostingResponse.AsObject;
  static serializeBinaryToWriter(message: CreateJobPostingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateJobPostingResponse;
  static deserializeBinaryFromReader(message: CreateJobPostingResponse, reader: jspb.BinaryReader): CreateJobPostingResponse;
}

export namespace CreateJobPostingResponse {
  export type AsObject = {
    jobPostingId: string,
  }
}

export class ReadJobPostingRequest extends jspb.Message {
  getJobPosting1(): string;
  setJobPosting1(value: string): ReadJobPostingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadJobPostingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadJobPostingRequest): ReadJobPostingRequest.AsObject;
  static serializeBinaryToWriter(message: ReadJobPostingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadJobPostingRequest;
  static deserializeBinaryFromReader(message: ReadJobPostingRequest, reader: jspb.BinaryReader): ReadJobPostingRequest;
}

export namespace ReadJobPostingRequest {
  export type AsObject = {
    jobPosting1: string,
  }
}

export class ReadJobPostingResponse extends jspb.Message {
  getIndustry(): string;
  setIndustry(value: string): ReadJobPostingResponse;

  getCompanyName(): string;
  setCompanyName(value: string): ReadJobPostingResponse;

  getLocation(): string;
  setLocation(value: string): ReadJobPostingResponse;

  getTitle(): string;
  setTitle(value: string): ReadJobPostingResponse;

  getBaseSalary(): Salary | undefined;
  setBaseSalary(value?: Salary): ReadJobPostingResponse;
  hasBaseSalary(): boolean;
  clearBaseSalary(): ReadJobPostingResponse;

  getBonusSalary(): Salary | undefined;
  setBonusSalary(value?: Salary): ReadJobPostingResponse;
  hasBonusSalary(): boolean;
  clearBonusSalary(): ReadJobPostingResponse;

  getQualificationsList(): Array<string>;
  setQualificationsList(value: Array<string>): ReadJobPostingResponse;
  clearQualificationsList(): ReadJobPostingResponse;
  addQualifications(value: string, index?: number): ReadJobPostingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadJobPostingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadJobPostingResponse): ReadJobPostingResponse.AsObject;
  static serializeBinaryToWriter(message: ReadJobPostingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadJobPostingResponse;
  static deserializeBinaryFromReader(message: ReadJobPostingResponse, reader: jspb.BinaryReader): ReadJobPostingResponse;
}

export namespace ReadJobPostingResponse {
  export type AsObject = {
    industry: string,
    companyName: string,
    location: string,
    title: string,
    baseSalary?: Salary.AsObject,
    bonusSalary?: Salary.AsObject,
    qualificationsList: Array<string>,
  }
}

export class UpdateJobPostingRequest extends jspb.Message {
  getJobPostingId(): string;
  setJobPostingId(value: string): UpdateJobPostingRequest;

  getIndustry(): string;
  setIndustry(value: string): UpdateJobPostingRequest;

  getCompanyName(): string;
  setCompanyName(value: string): UpdateJobPostingRequest;

  getLocation(): string;
  setLocation(value: string): UpdateJobPostingRequest;

  getTitle(): string;
  setTitle(value: string): UpdateJobPostingRequest;

  getBaseSalary(): Salary | undefined;
  setBaseSalary(value?: Salary): UpdateJobPostingRequest;
  hasBaseSalary(): boolean;
  clearBaseSalary(): UpdateJobPostingRequest;

  getBonusSalary(): Salary | undefined;
  setBonusSalary(value?: Salary): UpdateJobPostingRequest;
  hasBonusSalary(): boolean;
  clearBonusSalary(): UpdateJobPostingRequest;

  getQualificationsList(): Array<string>;
  setQualificationsList(value: Array<string>): UpdateJobPostingRequest;
  clearQualificationsList(): UpdateJobPostingRequest;
  addQualifications(value: string, index?: number): UpdateJobPostingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobPostingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobPostingRequest): UpdateJobPostingRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateJobPostingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobPostingRequest;
  static deserializeBinaryFromReader(message: UpdateJobPostingRequest, reader: jspb.BinaryReader): UpdateJobPostingRequest;
}

export namespace UpdateJobPostingRequest {
  export type AsObject = {
    jobPostingId: string,
    industry: string,
    companyName: string,
    location: string,
    title: string,
    baseSalary?: Salary.AsObject,
    bonusSalary?: Salary.AsObject,
    qualificationsList: Array<string>,
  }
}

export class UpdateJobPostingResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobPostingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobPostingResponse): UpdateJobPostingResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateJobPostingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobPostingResponse;
  static deserializeBinaryFromReader(message: UpdateJobPostingResponse, reader: jspb.BinaryReader): UpdateJobPostingResponse;
}

export namespace UpdateJobPostingResponse {
  export type AsObject = {
  }
}

export class DeleteJobPostingRequest extends jspb.Message {
  getJobPostingId(): string;
  setJobPostingId(value: string): DeleteJobPostingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteJobPostingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteJobPostingRequest): DeleteJobPostingRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteJobPostingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteJobPostingRequest;
  static deserializeBinaryFromReader(message: DeleteJobPostingRequest, reader: jspb.BinaryReader): DeleteJobPostingRequest;
}

export namespace DeleteJobPostingRequest {
  export type AsObject = {
    jobPostingId: string,
  }
}

export class DeleteJobPostingResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteJobPostingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteJobPostingResponse): DeleteJobPostingResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteJobPostingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteJobPostingResponse;
  static deserializeBinaryFromReader(message: DeleteJobPostingResponse, reader: jspb.BinaryReader): DeleteJobPostingResponse;
}

export namespace DeleteJobPostingResponse {
  export type AsObject = {
  }
}

export class Salary extends jspb.Message {
  getMinRange(): number;
  setMinRange(value: number): Salary;

  getMaxRange(): number;
  setMaxRange(value: number): Salary;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Salary.AsObject;
  static toObject(includeInstance: boolean, msg: Salary): Salary.AsObject;
  static serializeBinaryToWriter(message: Salary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Salary;
  static deserializeBinaryFromReader(message: Salary, reader: jspb.BinaryReader): Salary;
}

export namespace Salary {
  export type AsObject = {
    minRange: number,
    maxRange: number,
  }
}

