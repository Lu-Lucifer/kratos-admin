/**
 * 角色模块枚举映射常量
 */

type TFn = (key: string, options?: Record<string, any>) => string;

// ========== 角色状态（与菜单模块一致：ON/OFF） ==========

export const STATUS_COLORS: Record<string, string> = {
  ON: 'success',
  OFF: 'error',
};

export function getStatusMap(t: TFn) {
  return {
    ON: { text: t('statusMap.ON'), color: STATUS_COLORS.ON },
    OFF: { text: t('statusMap.OFF'), color: STATUS_COLORS.OFF },
  };
}

export function getStatusOptions(t: TFn) {
  return [
    { label: t('statusMap.ON'), value: 'ON' },
    { label: t('statusMap.OFF'), value: 'OFF' },
  ];
}

// ========== 权限树构建 ==========

interface TreeNode {
  key: number;
  title: string;
  children?: TreeNode[];
}

/**
 * 根据权限组和权限列表构建权限树
 * 权限组作为父节点，权限作为子节点
 */
export function buildPermissionTree(
  groups: Array<{ id?: number | string; title?: string; name?: string; code?: string }>,
  permissions: Array<{
    id?: number | string;
    title?: string;
    name?: string;
    code?: string;
    groupId?: number | string;
  }>,
): TreeNode[] {
  return (groups || []).map((group) => {
    const groupChildren = (permissions || [])
      .filter((p) => String(p.groupId) === String(group.id))
      .map((p) => ({
        key: Number(p.id),
        title: p.title || p.name || p.code || String(p.id),
      }));

    return {
      key: Number(group.id),
      title: group.title || group.name || group.code || String(group.id),
      children: groupChildren.length > 0 ? groupChildren : undefined,
    };
  });
}

/**
 * 从权限树勾选值中提取所有数字 ID（过滤掉非数字值）
 */
export function filterNumbers(values: any[]): number[] {
  if (!Array.isArray(values)) return [];
  return values
    .flat(Infinity)
    .filter((v) => typeof v === 'number' && !isNaN(v))
    .map((v) => Number(v));
}

/**
 * 递归收集树中所有「叶子节点」（无 children）的 key。
 * 用于提交勾选值时剥离父节点 key：
 * 这些 Tree 默认开启父子联动，勾选某父节点下全部子节点时父节点会被自动勾选，
 * 其 key（权限组/父菜单的 ID）会混入 checkedKeys。若直接 filterNumbers 提交，
 * 组/父菜单 ID 会被后端当作权限/菜单 ID 处理，可能造成越权绑定。
 * 这里用叶子集合对 checkedKeys 求交集，只保留真正的叶子 ID。
 */
export function extractLeafIds(checkedKeys: any[], treeData: any[]): number[] {
  if (!Array.isArray(checkedKeys) || !Array.isArray(treeData)) return [];
  const leafIds = new Set<number>();
  const collect = (nodes: any[]) => {
    for (const node of nodes) {
      if (node.children && node.children.length > 0) {
        collect(node.children);
      } else if (typeof node.key === 'number' && !isNaN(node.key)) {
        leafIds.add(node.key);
      }
    }
  };
  collect(treeData);
  return checkedKeys
    .flat(Infinity)
    .filter((v): v is number => typeof v === 'number' && !isNaN(v) && leafIds.has(v))
    .map((v) => Number(v));
}
