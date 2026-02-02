import { ExampleCodeGenerator, ExampleDataRender } from '../../../src/generator/exampleGenerator';
import { TestCodeModeler } from '@autorest/testmodeler/dist/src/core/model';
import { Helper } from '@autorest/testmodeler/dist/src/util/helper';
import { MockTestCodeGenerator, MockTestDataRender } from '../../../src/generator/mockTestGenerator';
import { processRequest } from '../../../src/generator/goTester';

describe('processRequest of go-tester', () => {
  let spyExampleRenderData;
  let spyExampleGenerateCode;
  let spyMockTestRenderData;
  let spyMockTestGenerateCode;

  beforeEach(() => {
    Helper.outputToModelerfour = jest.fn().mockResolvedValue(undefined);
    Helper.dump = jest.fn().mockResolvedValue(undefined);
    Helper.addCodeModelDump = jest.fn().mockResolvedValue(undefined);
    spyExampleRenderData = jest.spyOn(ExampleDataRender.prototype, 'renderData').mockReturnValue(undefined);
    spyExampleGenerateCode = jest.spyOn(ExampleCodeGenerator.prototype, 'generateCode').mockReturnValue(undefined);
    spyMockTestRenderData = jest.spyOn(MockTestDataRender.prototype, 'renderData').mockReturnValue(undefined);
    spyMockTestGenerateCode = jest.spyOn(MockTestCodeGenerator.prototype, 'generateCode').mockReturnValue(undefined);
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('processRequest with export-codemodel', async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-mock-test': true,
              'export-codemodel': true,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });

    await processRequest(undefined);

    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).toHaveBeenCalledTimes(1);
    expect(spyExampleRenderData).not.toHaveBeenCalled();
    expect(spyExampleGenerateCode).not.toHaveBeenCalled();
    expect(Helper.outputToModelerfour).toHaveBeenCalledTimes(1);
    expect(Helper.addCodeModelDump).toHaveBeenCalledTimes(2);
    expect(Helper.dump).toHaveBeenCalledTimes(1);
  });

  it('processRequest without export-codemodel', async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-mock-test': true,
              'export-codemodel': false,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });
    await processRequest(undefined);

    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).toHaveBeenCalledTimes(1);
    expect(spyExampleRenderData).not.toHaveBeenCalled();
    expect(spyExampleGenerateCode).not.toHaveBeenCalled();
    expect(Helper.outputToModelerfour).toHaveBeenCalledTimes(1);
    expect(Helper.addCodeModelDump).not.toHaveBeenCalled();
    expect(Helper.dump).toHaveBeenCalledTimes(1);
  });

  it("don't generate mock test if generate-mock-test is true", async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-mock-test': true,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });
    await processRequest(undefined);
    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).toHaveBeenCalledTimes(1);
    expect(spyExampleRenderData).not.toHaveBeenCalled();
    expect(spyExampleGenerateCode).not.toHaveBeenCalled();
  });

  it('generate sdk example if generate-sdk-example is true', async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-sdk-example': true,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });
    await processRequest(undefined);
    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).not.toHaveBeenCalled();
    expect(spyExampleRenderData).toHaveBeenCalledTimes(1);
    expect(spyExampleGenerateCode).toHaveBeenCalledTimes(1);
  });

  it('generate scenario test if generate-scenario-test is true', async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-scenario-test': true,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });
    await processRequest(undefined);
    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).not.toHaveBeenCalled();
    expect(spyExampleRenderData).not.toHaveBeenCalled();
    expect(spyExampleGenerateCode).not.toHaveBeenCalled();
  });

  it('generate sdk sample if generate-sdk-sample is true', async () => {
    TestCodeModeler.getSessionFromHost = jest.fn().mockResolvedValue({
      getValue: jest.fn().mockImplementation((key: string) => {
        if (key === '') {
          return {
            testmodeler: {
              'generate-sdk-sample': true,
            },
          };
        } else if (key === 'header-text') {
          return '';
        }
      }),
    });
    await processRequest(undefined);
    expect(spyMockTestRenderData).toHaveBeenCalledTimes(1);
    expect(spyMockTestGenerateCode).not.toHaveBeenCalled();
    expect(spyExampleRenderData).not.toHaveBeenCalled();
    expect(spyExampleGenerateCode).not.toHaveBeenCalled();
  });
});
