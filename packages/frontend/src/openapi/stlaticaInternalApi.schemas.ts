/**
 * Generated by orval v7.0.1 🍺
 * Do not edit manually.
 * stlatica_internal_api
 * stlatica internal api
 * OpenAPI spec version: 0.1.0
 */
export type LoginBody = {
  password?: string;
  preferred_user_id?: string;
};

export type UploadImage201 = {
  image_id?: string;
};

export type GetFollowersParams = {
  /**
 * userを識別するための一意のID。未指定の場合、先頭から取得

 */
  user_pagination_id?: UserPaginationIdParameter;
  /**
 * 取得数。default = 100

 */
  limit?: LimitParameter;
};

export type GetFollowUsersParams = {
  /**
 * userを識別するための一意のID。未指定の場合、先頭から取得

 */
  user_pagination_id?: UserPaginationIdParameter;
  /**
 * 取得数。default = 100

 */
  limit?: LimitParameter;
};

export type GetUsersParams = {
  /**
   * ユーザ名
   */
  user_name?: string;
};

export type CreateUserBody = {
  email?: string;
  name?: string;
};

export type TimelineToDateParameter = string;

export type TimelineFromDateParameter = string;

export type TimelineLimitParameter = number;

/**
 * timelineの種類
- home: 指定したuser_idのplatの配列
- following: 指定したuser_idのfollowingのplatの配列
- local: インスタンス内の全てのplatの配列

 */
export type TimelineTypeParameter =
  (typeof TimelineTypeParameter)[keyof typeof TimelineTypeParameter];

export const TimelineTypeParameter = {
  home: "home",
  following: "following",
  local: "local",
} as const;

export type TimelineUserIdParameter = string;

export type GetTimelineByQueryParams = {
  user_id: TimelineUserIdParameter;
  type: TimelineTypeParameter;
  limit?: TimelineLimitParameter;
  from_date?: TimelineFromDateParameter;
  to_date?: TimelineToDateParameter;
};

/**
 * 取得数。default = 100

 */
export type LimitParameter = number;

/**
 * userを識別するための一意のID。未指定の場合、先頭から取得

 */
export type UserPaginationIdParameter = string;

/**
 * service unavailable \
error code:
- 'SERVICE_UNAVAILABLE'

 */
export type E503Response = ErrorResponse;

/**
 * internal server error \
error code:
- 'INTERNAL_SERVER_ERROR'

 */
export type E500Response = ErrorResponse;

/**
 * not found \
error code:
- 'NOT_FOUND'

 */
export type E404Response = ErrorResponse;

/**
 * unauthorized \
error code:
- 'UNAUTHORIZED'

 */
export type E401Response = void;

/**
 * platを識別するための一意のID

 */
export type PlatID = string;

export type ErrorResponseCode = (typeof ErrorResponseCode)[keyof typeof ErrorResponseCode];

export const ErrorResponseCode = {
  BAD_REQUEST: "BAD_REQUEST",
  MISSING_PARAMETER: "MISSING_PARAMETER",
  UNAUTHORIZED: "UNAUTHORIZED",
  NOT_FOUND: "NOT_FOUND",
  INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
  SERVICE_UNAVAILABLE: "SERVICE_UNAVAILABLE",
  CONFLICT: "CONFLICT",
  UNPROCESSABLE_ENTITY: "UNPROCESSABLE_ENTITY",
} as const;

export interface ErrorResponse {
  code: ErrorResponseCode;
  message: string;
}

/**
 * bad request \
error code:
- 'BAD_REQUEST'
- 'MISSING_PARAMETER'

 */
export type E400Response = ErrorResponse;

/**
 * userを識別するための一意のID

 */
export type UserID = string;

/**
 * plat
 */
export interface PlatPost {
  /** platの本文 */
  content: string;
  user_id: UserID;
}

/**
 * plat
 */
export interface Plat {
  /** platの本文 */
  content: string;
  /** platが作成された日時(ISO8601) */
  created_at: string;
  /** platがお気に入りされた数 */
  favorite_count?: number;
  /** platに添付された画像のURL */
  image_urls?: string[];
  plat_id: PlatID;
  /** platがリプラットされた数 */
  replat_count?: number;
  /** platに対するリプライの数 */
  reply_count?: number;
  user_id?: UserID;
}

/**
 * user without follow counts
 */
export interface UserLightweight {
  /** ユーザのアイコン画像のURL */
  icon: string;
  /** 公開アカウントであるかどうか \
external apiのmanuallyApprovesFollowersと同一の値となる
 */
  is_public: boolean;
  /** ユーザのプロフィール */
  summary: string;
  user_id: UserID;
  /** 画面上に表示されるユーザ名 */
  username: string;
}

/**
 * user
 */
export interface User {
  /** follower数 */
  followers_count: number;
  /** following数 */
  following_count: number;
  /** ユーザのアイコン画像のURL */
  icon: string;
  /** 公開アカウントであるかどうか \
external apiのmanuallyApprovesFollowersと同一の値となる
 */
  is_public: boolean;
  /** ユーザのプロフィール */
  summary: string;
  user_id: UserID;
  /** 画面上に表示されるユーザ名 */
  username: string;
}
