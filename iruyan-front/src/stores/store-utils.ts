import type { StoreApi, UseBoundStore } from 'zustand';

/**
 * Zustandストア用のヘルパータイプとユーティリティ関数
 */

// ストアセレクター用の型安全なヘルパー
export type StoreSelector<T, U> = (state: T) => U;

// 浅い比較を使用するセレクター
export const createShallowSelector = <T, U>(
  selector: StoreSelector<T, U>
) => selector;

// ストアのアクションのみを抽出する型
export type StoreActions<T> = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [K in keyof T]: T[K] extends (...args: any[]) => any ? T[K] : never;
};

// ストアの状態のみを抽出する型
export type StoreState<T> = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [K in keyof T]: T[K] extends (...args: any[]) => any ? never : T[K];
};

// ストアの購読を管理するヘルパー
export const createStoreSubscription = <T>(
  store: UseBoundStore<StoreApi<T>>,
  listener: (state: T, prevState: T) => void
) => {
  return store.subscribe(listener);
};

// デバッグ用のストア状態ログ関数
export const logStoreState = <T>(
  storeName: string,
  state: T,
  action?: string
): void => {
  if (process.env.NODE_ENV === 'development') {
    console.group(`🗃️ Store: ${storeName}${action ? ` - ${action}` : ''}`);
    console.log('State:', state);
    console.groupEnd();
  }
};

// ストアのリセット機能を追加するための型
export interface ResettableStore {
  reset: () => void;
}