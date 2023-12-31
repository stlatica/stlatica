/**
 * Generated by orval v6.21.0 🍺
 * Do not edit manually.
 * stlatica_internal_api
 * stlatica internal api
 * OpenAPI spec version: 0.1.0
 */
import axios from 'axios'
import type {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
import useSwr from 'swr'
import type {
  Key,
  SWRConfiguration
} from 'swr'
import type {
  ErrorResponse,
  Plat,
  PlatPost,
  User
} from './stlaticaInternalApi.schemas'


  
  /**
 * get user
 * @summary get user
 */
export const getUser = (
    userId: string, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<User>> => {
    return axios.get(
      `http://localhost:4010/internal/v1/users/${userId}`,options
    );
  }


export const getGetUserKey = (userId: string,) => [`http://localhost:4010/internal/v1/users/${userId}`] as const;

    
export type GetUserQueryResult = NonNullable<Awaited<ReturnType<typeof getUser>>>
export type GetUserQueryError = AxiosError<ErrorResponse>

/**
 * @summary get user
 */
export const useGetUser = <TError = AxiosError<ErrorResponse>>(
 userId: string, options?: { swr?:SWRConfiguration<Awaited<ReturnType<typeof getUser>>, TError> & { swrKey?: Key, enabled?: boolean }, axios?: AxiosRequestConfig }

  ) => {

  const {swr: swrOptions, axios: axiosOptions} = options ?? {}

  const isEnabled = swrOptions?.enabled !== false && !!(userId)
    const swrKey = swrOptions?.swrKey ?? (() => isEnabled ? getGetUserKey(userId) : null);
  const swrFn = () => getUser(userId, axiosOptions);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query
  }
}


/**
 * Returns plat.
 * @summary Get plat by ID.
 */
export const getPlat = (
    platId: unknown, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<Plat>> => {
    return axios.get(
      `http://localhost:4010/internal/v1/plats/${platId}`,options
    );
  }


export const getGetPlatKey = (platId: unknown,) => [`http://localhost:4010/internal/v1/plats/${platId}`] as const;

    
export type GetPlatQueryResult = NonNullable<Awaited<ReturnType<typeof getPlat>>>
export type GetPlatQueryError = AxiosError<ErrorResponse | void>

/**
 * @summary Get plat by ID.
 */
export const useGetPlat = <TError = AxiosError<ErrorResponse | void>>(
 platId: unknown, options?: { swr?:SWRConfiguration<Awaited<ReturnType<typeof getPlat>>, TError> & { swrKey?: Key, enabled?: boolean }, axios?: AxiosRequestConfig }

  ) => {

  const {swr: swrOptions, axios: axiosOptions} = options ?? {}

  const isEnabled = swrOptions?.enabled !== false && !!(platId)
    const swrKey = swrOptions?.swrKey ?? (() => isEnabled ? getGetPlatKey(platId) : null);
  const swrFn = () => getPlat(platId, axiosOptions);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query
  }
}


/**
 * Delete plat.
 * @summary Delete plat by ID.
 */
export const deletePlat = (
    platId: unknown, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<void>> => {
    return axios.delete(
      `http://localhost:4010/internal/v1/plats/${platId}`,options
    );
  }



/**
 * Post plat.
 * @summary Post plat.
 */
export const postPlat = (
    platPost: PlatPost, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<void>> => {
    return axios.post(
      `http://localhost:4010/internal/v1/plats`,
      platPost,options
    );
  }



