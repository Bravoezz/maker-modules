//module test

//dir: models fileName: test.model.ts
export interface Test {}

//dir: contracts fileName: test.contract.ts
export interface ITestRepository {}
export interface ITestService {}

//dir: controllers fileName: test.controller.ts
export class TestController {}

//dir: services fileName: test.service.ts
// import {ITestService} from './contracts/test.contract.ts'
export class TestService implements ITestService {}

//dir: repositories fileName: test.repository.ts
// import {ITestRepository} from './contracts/test.contract.ts'
export class TestRepository implements ITestRepository {}