import { defineStore } from 'pinia';

import { defDepartmentService, makeQueryString, makeUpdateMask } from '#/rpc';

export const useDepartmentStore = defineStore('department', () => {
  /**
   * 查询部门列表
   */
  async function listDepartment(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDepartmentService.ListDepartment({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy: orderBy ?? [],
      query: makeQueryString(formValues ?? null),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 获取部门
   */
  async function getDepartment(id: number) {
    return await defDepartmentService.GetDepartment({ id });
  }

  /**
   * 创建部门
   */
  async function createDepartment(values: object) {
    return await defDepartmentService.CreateDepartment({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新部门
   */
  async function updateDepartment(id: number, values: object) {
    return await defDepartmentService.UpdateDepartment({
      data: {
        id,
        ...values,
        children: [],
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除部门
   */
  async function deleteDepartment(id: number) {
    return await defDepartmentService.DeleteDepartment({ id });
  }

  return {
    listDepartment,
    getDepartment,
    createDepartment,
    updateDepartment,
    deleteDepartment,
  };
});
