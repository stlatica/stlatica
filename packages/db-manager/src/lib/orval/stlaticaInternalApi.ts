/**
 * Generated by orval v7.1.0 🍺
 * Do not edit manually.
 * stlatica_internal_api
 * stlatica internal api
 * OpenAPI spec version: 0.1.0
 */
import { createMutation, createQuery } from '@tanstack/svelte-query';
import type {
	CreateMutationOptions,
	CreateMutationResult,
	CreateQueryOptions,
	CreateQueryResult,
	MutationFunction,
	QueryFunction,
	QueryKey
} from '@tanstack/svelte-query';
import axios from 'axios';
import type { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import type {
	CreateUserBody,
	E400Response,
	E401Response,
	E404Response,
	E409Response,
	E422Response,
	E500Response,
	E503Response,
	ErrorResponse,
	GetFollowersParams,
	GetFollowsParams,
	GetTimelineByQueryParams,
	GetUsersParams,
	LoginBody,
	Plat,
	PlatPost,
	UploadImage201,
	User,
	UserID,
	UserLightweight
} from './stlaticaInternalApi.schemas';

/**
 * get user
 * @summary get user
 */
export const getUser = (
	userId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<User>> => {
	return axios.get(`http://localhost:8080/internal/v1/users/${userId}`, options);
};

export const getGetUserQueryKey = (userId: string) => {
	return [`http://localhost:8080/internal/v1/users/${userId}`] as const;
};

export const getGetUserQueryOptions = <
	TData = Awaited<ReturnType<typeof getUser>>,
	TError = AxiosError<E400Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUser>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetUserQueryKey(userId);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getUser>>> = ({ signal }) =>
		getUser(userId, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!userId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getUser>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetUserQueryResult = NonNullable<Awaited<ReturnType<typeof getUser>>>;
export type GetUserQueryError = AxiosError<
	E400Response | E404Response | E500Response | E503Response
>;

/**
 * @summary get user
 */

export function createGetUser<
	TData = Awaited<ReturnType<typeof getUser>>,
	TError = AxiosError<E400Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUser>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetUserQueryOptions(userId, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Delete User
 * @summary Delete User by ID.
 */
export const deleteUser = (
	userId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.delete(`http://localhost:8080/internal/v1/users/${userId}`, options);
};

export const getDeleteUserMutationOptions = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteUser>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof deleteUser>>,
	TError,
	{ userId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteUser>>, { userId: string }> = (
		props
	) => {
		const { userId } = props ?? {};

		return deleteUser(userId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type DeleteUserMutationResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>;

export type DeleteUserMutationError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Delete User by ID.
 */
export const createDeleteUser = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteUser>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof deleteUser>>,
	TError,
	{ userId: string },
	TContext
> => {
	const mutationOptions = getDeleteUserMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * Get user icon.
 * @summary Get user icon.
 */
export const getUserIcon = (
	userId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<string>> => {
	return axios.get(`http://localhost:8080/internal/v1/users/${userId}/icon`, options);
};

export const getGetUserIconQueryKey = (userId: string) => {
	return [`http://localhost:8080/internal/v1/users/${userId}/icon`] as const;
};

export const getGetUserIconQueryOptions = <
	TData = Awaited<ReturnType<typeof getUserIcon>>,
	TError = AxiosError<E404Response | E500Response | E503Response>
>(
	userId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUserIcon>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetUserIconQueryKey(userId);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getUserIcon>>> = ({ signal }) =>
		getUserIcon(userId, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!userId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getUserIcon>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetUserIconQueryResult = NonNullable<Awaited<ReturnType<typeof getUserIcon>>>;
export type GetUserIconQueryError = AxiosError<E404Response | E500Response | E503Response>;

/**
 * @summary Get user icon.
 */

export function createGetUserIcon<
	TData = Awaited<ReturnType<typeof getUserIcon>>,
	TError = AxiosError<E404Response | E500Response | E503Response>
>(
	userId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUserIcon>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetUserIconQueryOptions(userId, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Create new user.
 * @summary Create new user.
 */
export const createUser = (
	createUserBody: CreateUserBody,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<UserID>> => {
	return axios.post(`http://localhost:8080/internal/v1/users`, createUserBody, options);
};

export const getCreateUserMutationOptions = <
	TError = AxiosError<
		E400Response | E404Response | E409Response | E422Response | E500Response | E503Response
	>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof createUser>>,
		TError,
		{ data: CreateUserBody },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof createUser>>,
	TError,
	{ data: CreateUserBody },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<
		Awaited<ReturnType<typeof createUser>>,
		{ data: CreateUserBody }
	> = (props) => {
		const { data } = props ?? {};

		return createUser(data, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type CreateUserMutationResult = NonNullable<Awaited<ReturnType<typeof createUser>>>;
export type CreateUserMutationBody = CreateUserBody;
export type CreateUserMutationError = AxiosError<
	E400Response | E404Response | E409Response | E422Response | E500Response | E503Response
>;

/**
 * @summary Create new user.
 */
export const createCreateUser = <
	TError = AxiosError<
		E400Response | E404Response | E409Response | E422Response | E500Response | E503Response
	>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof createUser>>,
		TError,
		{ data: CreateUserBody },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof createUser>>,
	TError,
	{ data: CreateUserBody },
	TContext
> => {
	const mutationOptions = getCreateUserMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * Get users.
 * @summary Get users.
 */
export const getUsers = (
	params?: GetUsersParams,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<User[]>> => {
	return axios.get(`http://localhost:8080/internal/v1/users`, {
		...options,
		params: { ...params, ...options?.params }
	});
};

export const getGetUsersQueryKey = (params?: GetUsersParams) => {
	return [`http://localhost:8080/internal/v1/users`, ...(params ? [params] : [])] as const;
};

export const getGetUsersQueryOptions = <
	TData = Awaited<ReturnType<typeof getUsers>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	params?: GetUsersParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUsers>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetUsersQueryKey(params);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getUsers>>> = ({ signal }) =>
		getUsers(params, { signal, ...axiosOptions });

	return { queryKey, queryFn, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getUsers>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetUsersQueryResult = NonNullable<Awaited<ReturnType<typeof getUsers>>>;
export type GetUsersQueryError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Get users.
 */

export function createGetUsers<
	TData = Awaited<ReturnType<typeof getUsers>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	params?: GetUsersParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getUsers>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetUsersQueryOptions(params, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Get follow user list.
 * @summary Get follow user list.
 */
export const getFollows = (
	userId: string,
	params?: GetFollowsParams,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<UserLightweight[]>> => {
	return axios.get(`http://localhost:8080/internal/v1/users/${userId}/follows`, {
		...options,
		params: { ...params, ...options?.params }
	});
};

export const getGetFollowsQueryKey = (userId: string, params?: GetFollowsParams) => {
	return [
		`http://localhost:8080/internal/v1/users/${userId}/follows`,
		...(params ? [params] : [])
	] as const;
};

export const getGetFollowsQueryOptions = <
	TData = Awaited<ReturnType<typeof getFollows>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	params?: GetFollowsParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getFollows>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetFollowsQueryKey(userId, params);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getFollows>>> = ({ signal }) =>
		getFollows(userId, params, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!userId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getFollows>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetFollowsQueryResult = NonNullable<Awaited<ReturnType<typeof getFollows>>>;
export type GetFollowsQueryError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Get follow user list.
 */

export function createGetFollows<
	TData = Awaited<ReturnType<typeof getFollows>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	params?: GetFollowsParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getFollows>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetFollowsQueryOptions(userId, params, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Get follower list.
 * @summary Get follower list.
 */
export const getFollowers = (
	userId: string,
	params?: GetFollowersParams,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<UserLightweight[]>> => {
	return axios.get(`http://localhost:8080/internal/v1/users/${userId}/followers`, {
		...options,
		params: { ...params, ...options?.params }
	});
};

export const getGetFollowersQueryKey = (userId: string, params?: GetFollowersParams) => {
	return [
		`http://localhost:8080/internal/v1/users/${userId}/followers`,
		...(params ? [params] : [])
	] as const;
};

export const getGetFollowersQueryOptions = <
	TData = Awaited<ReturnType<typeof getFollowers>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	params?: GetFollowersParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getFollowers>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetFollowersQueryKey(userId, params);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getFollowers>>> = ({ signal }) =>
		getFollowers(userId, params, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!userId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getFollowers>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetFollowersQueryResult = NonNullable<Awaited<ReturnType<typeof getFollowers>>>;
export type GetFollowersQueryError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Get follower list.
 */

export function createGetFollowers<
	TData = Awaited<ReturnType<typeof getFollowers>>,
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>
>(
	userId: string,
	params?: GetFollowersParams,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getFollowers>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetFollowersQueryOptions(userId, params, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Add new follow user.
 * @summary Add new follow user.
 */
export const postFollow = (
	userId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.post(`http://localhost:8080/internal/v1/users/${userId}/follow`, undefined, options);
};

export const getPostFollowMutationOptions = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postFollow>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof postFollow>>,
	TError,
	{ userId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof postFollow>>, { userId: string }> = (
		props
	) => {
		const { userId } = props ?? {};

		return postFollow(userId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type PostFollowMutationResult = NonNullable<Awaited<ReturnType<typeof postFollow>>>;

export type PostFollowMutationError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Add new follow user.
 */
export const createPostFollow = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postFollow>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof postFollow>>,
	TError,
	{ userId: string },
	TContext
> => {
	const mutationOptions = getPostFollowMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * Delete follow user.
 * @summary Delete follow user.
 */
export const deleteFollow = (
	userId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.delete(`http://localhost:8080/internal/v1/users/${userId}/follow`, options);
};

export const getDeleteFollowMutationOptions = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteFollow>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof deleteFollow>>,
	TError,
	{ userId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<
		Awaited<ReturnType<typeof deleteFollow>>,
		{ userId: string }
	> = (props) => {
		const { userId } = props ?? {};

		return deleteFollow(userId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type DeleteFollowMutationResult = NonNullable<Awaited<ReturnType<typeof deleteFollow>>>;

export type DeleteFollowMutationError = AxiosError<
	E400Response | E401Response | E404Response | E500Response | E503Response
>;

/**
 * @summary Delete follow user.
 */
export const createDeleteFollow = <
	TError = AxiosError<E400Response | E401Response | E404Response | E500Response | E503Response>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteFollow>>,
		TError,
		{ userId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof deleteFollow>>,
	TError,
	{ userId: string },
	TContext
> => {
	const mutationOptions = getDeleteFollowMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * Returns plat.
 * @summary Get plat by ID.
 */
export const getPlat = (
	platId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<Plat>> => {
	return axios.get(`http://localhost:8080/internal/v1/plats/${platId}`, options);
};

export const getGetPlatQueryKey = (platId: string) => {
	return [`http://localhost:8080/internal/v1/plats/${platId}`] as const;
};

export const getGetPlatQueryOptions = <
	TData = Awaited<ReturnType<typeof getPlat>>,
	TError = AxiosError<ErrorResponse | void>
>(
	platId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getPlat>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetPlatQueryKey(platId);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getPlat>>> = ({ signal }) =>
		getPlat(platId, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!platId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getPlat>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetPlatQueryResult = NonNullable<Awaited<ReturnType<typeof getPlat>>>;
export type GetPlatQueryError = AxiosError<ErrorResponse | void>;

/**
 * @summary Get plat by ID.
 */

export function createGetPlat<
	TData = Awaited<ReturnType<typeof getPlat>>,
	TError = AxiosError<ErrorResponse | void>
>(
	platId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getPlat>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetPlatQueryOptions(platId, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * Delete plat.
 * @summary Delete plat by ID.
 */
export const deletePlat = (
	platId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.delete(`http://localhost:8080/internal/v1/plats/${platId}`, options);
};

export const getDeletePlatMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deletePlat>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof deletePlat>>,
	TError,
	{ platId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof deletePlat>>, { platId: string }> = (
		props
	) => {
		const { platId } = props ?? {};

		return deletePlat(platId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type DeletePlatMutationResult = NonNullable<Awaited<ReturnType<typeof deletePlat>>>;

export type DeletePlatMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary Delete plat by ID.
 */
export const createDeletePlat = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deletePlat>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof deletePlat>>,
	TError,
	{ platId: string },
	TContext
> => {
	const mutationOptions = getDeletePlatMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * @summary Add "favorite" to a specific plat by ID.
 */
export const postFavorite = (
	platId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.post(
		`http://localhost:8080/internal/v1/plats/${platId}/favorites`,
		undefined,
		options
	);
};

export const getPostFavoriteMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postFavorite>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof postFavorite>>,
	TError,
	{ platId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<
		Awaited<ReturnType<typeof postFavorite>>,
		{ platId: string }
	> = (props) => {
		const { platId } = props ?? {};

		return postFavorite(platId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type PostFavoriteMutationResult = NonNullable<Awaited<ReturnType<typeof postFavorite>>>;

export type PostFavoriteMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary Add "favorite" to a specific plat by ID.
 */
export const createPostFavorite = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postFavorite>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof postFavorite>>,
	TError,
	{ platId: string },
	TContext
> => {
	const mutationOptions = getPostFavoriteMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * @summary Delete "favorite" to a specific plat by ID.
 */
export const deleteFavorite = (
	platId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.delete(`http://localhost:8080/internal/v1/plats/${platId}/favorites`, options);
};

export const getDeleteFavoriteMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteFavorite>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof deleteFavorite>>,
	TError,
	{ platId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<
		Awaited<ReturnType<typeof deleteFavorite>>,
		{ platId: string }
	> = (props) => {
		const { platId } = props ?? {};

		return deleteFavorite(platId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type DeleteFavoriteMutationResult = NonNullable<Awaited<ReturnType<typeof deleteFavorite>>>;

export type DeleteFavoriteMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary Delete "favorite" to a specific plat by ID.
 */
export const createDeleteFavorite = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteFavorite>>,
		TError,
		{ platId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof deleteFavorite>>,
	TError,
	{ platId: string },
	TContext
> => {
	const mutationOptions = getDeleteFavoriteMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * Post plat.
 * @summary Post plat.
 */
export const postPlat = (
	platPost: PlatPost,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.post(`http://localhost:8080/internal/v1/plats`, platPost, options);
};

export const getPostPlatMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postPlat>>,
		TError,
		{ data: PlatPost },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof postPlat>>,
	TError,
	{ data: PlatPost },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof postPlat>>, { data: PlatPost }> = (
		props
	) => {
		const { data } = props ?? {};

		return postPlat(data, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type PostPlatMutationResult = NonNullable<Awaited<ReturnType<typeof postPlat>>>;
export type PostPlatMutationBody = PlatPost;
export type PostPlatMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary Post plat.
 */
export const createPostPlat = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof postPlat>>,
		TError,
		{ data: PlatPost },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof postPlat>>,
	TError,
	{ data: PlatPost },
	TContext
> => {
	const mutationOptions = getPostPlatMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * get timeline \
plat idの降順でplatの配列を取得する

 * @summary get timeline
 */
export const getTimeline = (
	timelineId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<Plat[]>> => {
	return axios.get(`http://localhost:8080/internal/v1/timelines/${timelineId}`, options);
};

export const getGetTimelineQueryKey = (timelineId: string) => {
	return [`http://localhost:8080/internal/v1/timelines/${timelineId}`] as const;
};

export const getGetTimelineQueryOptions = <
	TData = Awaited<ReturnType<typeof getTimeline>>,
	TError = AxiosError<ErrorResponse>
>(
	timelineId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getTimeline>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetTimelineQueryKey(timelineId);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getTimeline>>> = ({ signal }) =>
		getTimeline(timelineId, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!timelineId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getTimeline>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetTimelineQueryResult = NonNullable<Awaited<ReturnType<typeof getTimeline>>>;
export type GetTimelineQueryError = AxiosError<ErrorResponse>;

/**
 * @summary get timeline
 */

export function createGetTimeline<
	TData = Awaited<ReturnType<typeof getTimeline>>,
	TError = AxiosError<ErrorResponse>
>(
	timelineId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getTimeline>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetTimelineQueryOptions(timelineId, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * get timeline by query params \
指定したクエリパラメータに一致するplatの配列を取得する

 * @summary get timeline by query params
 */
export const getTimelineByQuery = (
	params: GetTimelineByQueryParams,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<Plat[]>> => {
	return axios.get(`http://stlatica_server:8080/internal/v1/timelines`, {
		...options,
		params: { ...params, ...options?.params }
	});
};

export const getGetTimelineByQueryQueryKey = (params: GetTimelineByQueryParams) => {
	return [`http://localhost:8080/internal/v1/timelines`, ...(params ? [params] : [])] as const;
};

export const getGetTimelineByQueryQueryOptions = <
	TData = Awaited<ReturnType<typeof getTimelineByQuery>>,
	TError = AxiosError<ErrorResponse>
>(
	params: GetTimelineByQueryParams,
	options?: {
		query?: Partial<
			CreateQueryOptions<Awaited<ReturnType<typeof getTimelineByQuery>>, TError, TData>
		>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetTimelineByQueryQueryKey(params);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getTimelineByQuery>>> = ({ signal }) =>
		getTimelineByQuery(params, { signal, ...axiosOptions });

	return { queryKey, queryFn, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getTimelineByQuery>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetTimelineByQueryQueryResult = NonNullable<
	Awaited<ReturnType<typeof getTimelineByQuery>>
>;
export type GetTimelineByQueryQueryError = AxiosError<ErrorResponse>;

/**
 * @summary get timeline by query params
 */

export function createGetTimelineByQuery<
	TData = Awaited<ReturnType<typeof getTimelineByQuery>>,
	TError = AxiosError<ErrorResponse>
>(
	params: GetTimelineByQueryParams,
	options?: {
		query?: Partial<
			CreateQueryOptions<Awaited<ReturnType<typeof getTimelineByQuery>>, TError, TData>
		>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetTimelineByQueryQueryOptions(params, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * get image \
base64形式で返される

 * @summary get image
 */
export const getImage = (
	imageId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<string>> => {
	return axios.get(`http://localhost:8080/internal/v1/images/${imageId}`, options);
};

export const getGetImageQueryKey = (imageId: string) => {
	return [`http://localhost:8080/internal/v1/images/${imageId}`] as const;
};

export const getGetImageQueryOptions = <
	TData = Awaited<ReturnType<typeof getImage>>,
	TError = AxiosError<ErrorResponse | void>
>(
	imageId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getImage>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
) => {
	const { query: queryOptions, axios: axiosOptions } = options ?? {};

	const queryKey = queryOptions?.queryKey ?? getGetImageQueryKey(imageId);

	const queryFn: QueryFunction<Awaited<ReturnType<typeof getImage>>> = ({ signal }) =>
		getImage(imageId, { signal, ...axiosOptions });

	return { queryKey, queryFn, enabled: !!imageId, ...queryOptions } as CreateQueryOptions<
		Awaited<ReturnType<typeof getImage>>,
		TError,
		TData
	> & { queryKey: QueryKey };
};

export type GetImageQueryResult = NonNullable<Awaited<ReturnType<typeof getImage>>>;
export type GetImageQueryError = AxiosError<ErrorResponse | void>;

/**
 * @summary get image
 */

export function createGetImage<
	TData = Awaited<ReturnType<typeof getImage>>,
	TError = AxiosError<ErrorResponse | void>
>(
	imageId: string,
	options?: {
		query?: Partial<CreateQueryOptions<Awaited<ReturnType<typeof getImage>>, TError, TData>>;
		axios?: AxiosRequestConfig;
	}
): CreateQueryResult<TData, TError> & { queryKey: QueryKey } {
	const queryOptions = getGetImageQueryOptions(imageId, options);

	const query = createQuery(queryOptions) as CreateQueryResult<TData, TError> & {
		queryKey: QueryKey;
	};

	query.queryKey = queryOptions.queryKey;

	return query;
}

/**
 * delete image \
ステータスコードのみを返す

 * @summary delete image
 */
export const deleteImage = (
	imageId: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.delete(`http://localhost:8080/internal/v1/images/${imageId}`, options);
};

export const getDeleteImageMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteImage>>,
		TError,
		{ imageId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof deleteImage>>,
	TError,
	{ imageId: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<
		Awaited<ReturnType<typeof deleteImage>>,
		{ imageId: string }
	> = (props) => {
		const { imageId } = props ?? {};

		return deleteImage(imageId, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type DeleteImageMutationResult = NonNullable<Awaited<ReturnType<typeof deleteImage>>>;

export type DeleteImageMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary delete image
 */
export const createDeleteImage = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof deleteImage>>,
		TError,
		{ imageId: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof deleteImage>>,
	TError,
	{ imageId: string },
	TContext
> => {
	const mutationOptions = getDeleteImageMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * upload image \
base64形式で受け取る

 * @summary upload image
 */
export const uploadImage = (
	uploadImageBody: string,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<UploadImage201>> => {
	return axios.post(`http://localhost:8080/internal/v1/images`, uploadImageBody, options);
};

export const getUploadImageMutationOptions = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof uploadImage>>,
		TError,
		{ data: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof uploadImage>>,
	TError,
	{ data: string },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof uploadImage>>, { data: string }> = (
		props
	) => {
		const { data } = props ?? {};

		return uploadImage(data, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type UploadImageMutationResult = NonNullable<Awaited<ReturnType<typeof uploadImage>>>;
export type UploadImageMutationBody = string;
export type UploadImageMutationError = AxiosError<ErrorResponse | void>;

/**
 * @summary upload image
 */
export const createUploadImage = <
	TError = AxiosError<ErrorResponse | void>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof uploadImage>>,
		TError,
		{ data: string },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof uploadImage>>,
	TError,
	{ data: string },
	TContext
> => {
	const mutationOptions = getUploadImageMutationOptions(options);

	return createMutation(mutationOptions);
};

/**
 * login
 * @summary login
 */
export const login = (
	loginBody: LoginBody,
	options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
	return axios.post(`http://localhost:8080/internal/v1/login`, loginBody, options);
};

export const getLoginMutationOptions = <
	TError = AxiosError<ErrorResponse>,
	TContext = unknown
>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof login>>,
		TError,
		{ data: LoginBody },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationOptions<
	Awaited<ReturnType<typeof login>>,
	TError,
	{ data: LoginBody },
	TContext
> => {
	const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

	const mutationFn: MutationFunction<Awaited<ReturnType<typeof login>>, { data: LoginBody }> = (
		props
	) => {
		const { data } = props ?? {};

		return login(data, axiosOptions);
	};

	return { mutationFn, ...mutationOptions };
};

export type LoginMutationResult = NonNullable<Awaited<ReturnType<typeof login>>>;
export type LoginMutationBody = LoginBody;
export type LoginMutationError = AxiosError<ErrorResponse>;

/**
 * @summary login
 */
export const createLogin = <TError = AxiosError<ErrorResponse>, TContext = unknown>(options?: {
	mutation?: CreateMutationOptions<
		Awaited<ReturnType<typeof login>>,
		TError,
		{ data: LoginBody },
		TContext
	>;
	axios?: AxiosRequestConfig;
}): CreateMutationResult<
	Awaited<ReturnType<typeof login>>,
	TError,
	{ data: LoginBody },
	TContext
> => {
	const mutationOptions = getLoginMutationOptions(options);

	return createMutation(mutationOptions);
};
