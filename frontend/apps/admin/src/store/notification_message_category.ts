import { defineStore } from 'pinia';

import {
  defNotificationMessageCategoryService,
  makeQueryString,
  makeUpdateMask,
} from '#/rpc';

export const useNotificationMessageCategoryStore = defineStore(
  'notification_message_category',
  () => {
    /**
     * 查询通知消息列表
     */
    async function listNotificationMessageCategory(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await defNotificationMessageCategoryService.ListNotificationMessageCategory(
        {
          // @ts-ignore proto generated code is error.
          fieldMask,
          orderBy: orderBy ?? [],
          query: makeQueryString(formValues ?? null),
          page,
          pageSize,
          noPaging,
        },
      );
    }

    /**
     * 获取通知消息
     */
    async function getNotificationMessageCategory(id: number) {
      return await defNotificationMessageCategoryService.GetNotificationMessageCategory(
        { id },
      );
    }

    /**
     * 创建通知消息
     */
    async function createNotificationMessageCategory(values: object) {
      return await defNotificationMessageCategoryService.CreateNotificationMessageCategory(
        {
          // @ts-ignore proto generated code is error.
          data: {
            ...values,
          },
        },
      );
    }

    /**
     * 更新通知消息
     */
    async function updateNotificationMessageCategory(
      id: number,
      values: object,
    ) {
      return await defNotificationMessageCategoryService.UpdateNotificationMessageCategory(
        {
          // @ts-ignore proto generated code is error.
          data: {
            id,
            ...values,
          },
          // @ts-ignore proto generated code is error.
          updateMask: makeUpdateMask(Object.keys(values ?? [])),
        },
      );
    }

    /**
     * 删除通知消息
     */
    async function deleteNotificationMessageCategory(id: number) {
      return await defNotificationMessageCategoryService.DeleteNotificationMessageCategory(
        {
          id,
        },
      );
    }

    function $reset() {}

    return {
      $reset,
      listNotificationMessageCategory,
      getNotificationMessageCategory,
      createNotificationMessageCategory,
      updateNotificationMessageCategory,
      deleteNotificationMessageCategory,
    };
  },
);
