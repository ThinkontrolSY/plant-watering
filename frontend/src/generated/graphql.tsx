import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Mutation = {
  __typename?: 'Mutation';
  water: Scalars['Boolean']['output'];
};


export type MutationWaterArgs = {
  input: WateringInput;
};

export type Query = {
  __typename?: 'Query';
  channels: Array<Scalars['String']['output']>;
  waterStatistic?: Maybe<WaterStatistic>;
  weather?: Maybe<Weather>;
};


export type QueryWaterStatisticArgs = {
  channel: Scalars['String']['input'];
};

export type WateringInput = {
  channel: Scalars['String']['input'];
  seconds: Scalars['Int']['input'];
};

export type Weather = {
  __typename?: 'Weather';
  dayTemperature: Scalars['Float']['output'];
  nightTemperature: Scalars['Float']['output'];
  waterPlanSec: Scalars['Int']['output'];
  weather: Scalars['String']['output'];
  windDirection: Scalars['String']['output'];
  windPower: Scalars['String']['output'];
};

export type WaterStatistic = {
  __typename?: 'waterStatistic';
  autoWatering: Scalars['Int']['output'];
  manualWatering: Scalars['Int']['output'];
};

export type ChannelsQueryVariables = Exact<{ [key: string]: never; }>;


export type ChannelsQuery = { __typename?: 'Query', channels: Array<string> };

export type WeatherQueryVariables = Exact<{ [key: string]: never; }>;


export type WeatherQuery = { __typename?: 'Query', weather?: { __typename?: 'Weather', dayTemperature: number, nightTemperature: number, windDirection: string, windPower: string, weather: string, waterPlanSec: number } | null };

export type WaterStatisticQueryVariables = Exact<{
  channel: Scalars['String']['input'];
}>;


export type WaterStatisticQuery = { __typename?: 'Query', waterStatistic?: { __typename?: 'waterStatistic', autoWatering: number, manualWatering: number } | null };

export type WaterMutationVariables = Exact<{
  input: WateringInput;
}>;


export type WaterMutation = { __typename?: 'Mutation', water: boolean };


export const ChannelsDocument = gql`
    query channels {
  channels
}
    `;

/**
 * __useChannelsQuery__
 *
 * To run a query within a React component, call `useChannelsQuery` and pass it any options that fit your needs.
 * When your component renders, `useChannelsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useChannelsQuery({
 *   variables: {
 *   },
 * });
 */
export function useChannelsQuery(baseOptions?: Apollo.QueryHookOptions<ChannelsQuery, ChannelsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ChannelsQuery, ChannelsQueryVariables>(ChannelsDocument, options);
      }
export function useChannelsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ChannelsQuery, ChannelsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ChannelsQuery, ChannelsQueryVariables>(ChannelsDocument, options);
        }
export function useChannelsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ChannelsQuery, ChannelsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ChannelsQuery, ChannelsQueryVariables>(ChannelsDocument, options);
        }
export type ChannelsQueryHookResult = ReturnType<typeof useChannelsQuery>;
export type ChannelsLazyQueryHookResult = ReturnType<typeof useChannelsLazyQuery>;
export type ChannelsSuspenseQueryHookResult = ReturnType<typeof useChannelsSuspenseQuery>;
export type ChannelsQueryResult = Apollo.QueryResult<ChannelsQuery, ChannelsQueryVariables>;
export function refetchChannelsQuery(variables?: ChannelsQueryVariables) {
      return { query: ChannelsDocument, variables: variables }
    }
export const WeatherDocument = gql`
    query weather {
  weather {
    dayTemperature
    nightTemperature
    windDirection
    windPower
    weather
    waterPlanSec
  }
}
    `;

/**
 * __useWeatherQuery__
 *
 * To run a query within a React component, call `useWeatherQuery` and pass it any options that fit your needs.
 * When your component renders, `useWeatherQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useWeatherQuery({
 *   variables: {
 *   },
 * });
 */
export function useWeatherQuery(baseOptions?: Apollo.QueryHookOptions<WeatherQuery, WeatherQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<WeatherQuery, WeatherQueryVariables>(WeatherDocument, options);
      }
export function useWeatherLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<WeatherQuery, WeatherQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<WeatherQuery, WeatherQueryVariables>(WeatherDocument, options);
        }
export function useWeatherSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<WeatherQuery, WeatherQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<WeatherQuery, WeatherQueryVariables>(WeatherDocument, options);
        }
export type WeatherQueryHookResult = ReturnType<typeof useWeatherQuery>;
export type WeatherLazyQueryHookResult = ReturnType<typeof useWeatherLazyQuery>;
export type WeatherSuspenseQueryHookResult = ReturnType<typeof useWeatherSuspenseQuery>;
export type WeatherQueryResult = Apollo.QueryResult<WeatherQuery, WeatherQueryVariables>;
export function refetchWeatherQuery(variables?: WeatherQueryVariables) {
      return { query: WeatherDocument, variables: variables }
    }
export const WaterStatisticDocument = gql`
    query waterStatistic($channel: String!) {
  waterStatistic(channel: $channel) {
    autoWatering
    manualWatering
  }
}
    `;

/**
 * __useWaterStatisticQuery__
 *
 * To run a query within a React component, call `useWaterStatisticQuery` and pass it any options that fit your needs.
 * When your component renders, `useWaterStatisticQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useWaterStatisticQuery({
 *   variables: {
 *      channel: // value for 'channel'
 *   },
 * });
 */
export function useWaterStatisticQuery(baseOptions: Apollo.QueryHookOptions<WaterStatisticQuery, WaterStatisticQueryVariables> & ({ variables: WaterStatisticQueryVariables; skip?: boolean; } | { skip: boolean; }) ) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<WaterStatisticQuery, WaterStatisticQueryVariables>(WaterStatisticDocument, options);
      }
export function useWaterStatisticLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<WaterStatisticQuery, WaterStatisticQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<WaterStatisticQuery, WaterStatisticQueryVariables>(WaterStatisticDocument, options);
        }
export function useWaterStatisticSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<WaterStatisticQuery, WaterStatisticQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<WaterStatisticQuery, WaterStatisticQueryVariables>(WaterStatisticDocument, options);
        }
export type WaterStatisticQueryHookResult = ReturnType<typeof useWaterStatisticQuery>;
export type WaterStatisticLazyQueryHookResult = ReturnType<typeof useWaterStatisticLazyQuery>;
export type WaterStatisticSuspenseQueryHookResult = ReturnType<typeof useWaterStatisticSuspenseQuery>;
export type WaterStatisticQueryResult = Apollo.QueryResult<WaterStatisticQuery, WaterStatisticQueryVariables>;
export function refetchWaterStatisticQuery(variables: WaterStatisticQueryVariables) {
      return { query: WaterStatisticDocument, variables: variables }
    }
export const WaterDocument = gql`
    mutation water($input: WateringInput!) {
  water(input: $input)
}
    `;
export type WaterMutationFn = Apollo.MutationFunction<WaterMutation, WaterMutationVariables>;

/**
 * __useWaterMutation__
 *
 * To run a mutation, you first call `useWaterMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useWaterMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [waterMutation, { data, loading, error }] = useWaterMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useWaterMutation(baseOptions?: Apollo.MutationHookOptions<WaterMutation, WaterMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<WaterMutation, WaterMutationVariables>(WaterDocument, options);
      }
export type WaterMutationHookResult = ReturnType<typeof useWaterMutation>;
export type WaterMutationResult = Apollo.MutationResult<WaterMutation>;
export type WaterMutationOptions = Apollo.BaseMutationOptions<WaterMutation, WaterMutationVariables>;