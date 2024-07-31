/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  DateTime: { input: any; output: any; }
  Long: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type BetChinchiroGameTurnParticipant = {
  betAmount: Scalars['Int']['input'];
  turnId: Scalars['ID']['input'];
};

export type BetChinchiroGameTurnParticipantPayload = {
  __typename?: 'BetChinchiroGameTurnParticipantPayload';
  ok: Scalars['Boolean']['output'];
};

export enum ChinchiroCombination {
  Gonome = 'Gonome',
  Gozoro = 'Gozoro',
  Hifumi = 'Hifumi',
  Ichinome = 'Ichinome',
  Menashi = 'Menashi',
  Ninome = 'Ninome',
  Nizoro = 'Nizoro',
  Pinzoro = 'Pinzoro',
  Rokunome = 'Rokunome',
  Rokuzoro = 'Rokuzoro',
  Sannome = 'Sannome',
  Sanzoro = 'Sanzoro',
  Shigoro = 'Shigoro',
  Yonnome = 'Yonnome',
  Yonzoro = 'Yonzoro'
}

export type ChinchiroGame = {
  __typename?: 'ChinchiroGame';
  id: Scalars['ID']['output'];
  participants: Array<ChinchiroGameParticipant>;
  status: ChinchiroGameStatus;
  turns: Array<ChinchiroGameTurn>;
};

export type ChinchiroGameParticipant = {
  __typename?: 'ChinchiroGameParticipant';
  balance: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  roomParticipant: ChinchiroRoomParticipant;
  turnOrder: Scalars['Int']['output'];
};

export enum ChinchiroGameStatus {
  Finished = 'Finished',
  Progress = 'Progress'
}

export type ChinchiroGameTurn = {
  __typename?: 'ChinchiroGameTurn';
  dealer: ChinchiroGameParticipant;
  id: Scalars['ID']['output'];
  nextRoller?: Maybe<ChinchiroGameParticipant>;
  results: Array<ChinchiroGameTurnParticipantResult>;
  rolls: Array<ChinchiroGameTurnParticipantRoll>;
  status: ChinchiroGameTurnStatus;
  turnNumber: Scalars['Int']['output'];
};

export type ChinchiroGameTurnParticipantResult = {
  __typename?: 'ChinchiroGameTurnParticipantResult';
  combination: ChinchiroCombination;
  dice1: Scalars['Int']['output'];
  dice2: Scalars['Int']['output'];
  dice3: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  participant: ChinchiroGameParticipant;
  turn: ChinchiroGameTurn;
  winnings: Scalars['Int']['output'];
};

export type ChinchiroGameTurnParticipantResultsQuery = {
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  paging?: InputMaybe<PageableQuery>;
  turnId?: InputMaybe<Scalars['ID']['input']>;
};

export type ChinchiroGameTurnParticipantRoll = {
  __typename?: 'ChinchiroGameTurnParticipantRoll';
  dice1: Scalars['Int']['output'];
  dice2: Scalars['Int']['output'];
  dice3: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  participant: ChinchiroGameParticipant;
  rollNumber: Scalars['Int']['output'];
  turn: ChinchiroGameTurn;
};

export type ChinchiroGameTurnRollsQuery = {
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  paging?: InputMaybe<PageableQuery>;
  turnId?: InputMaybe<Scalars['ID']['input']>;
};

export enum ChinchiroGameTurnStatus {
  Betting = 'Betting',
  Finished = 'Finished',
  Rolling = 'Rolling'
}

export type ChinchiroGameTurnsQuery = {
  gameId?: InputMaybe<Scalars['ID']['input']>;
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  paging?: InputMaybe<PageableQuery>;
  statuses?: InputMaybe<Array<ChinchiroGameTurnStatus>>;
};

export type ChinchiroGamesQuery = {
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  name?: InputMaybe<Scalars['String']['input']>;
  paging?: InputMaybe<PageableQuery>;
  roomId?: InputMaybe<Scalars['ID']['input']>;
  statuses?: InputMaybe<Array<ChinchiroGameStatus>>;
};

export type ChinchiroRoom = {
  __typename?: 'ChinchiroRoom';
  games: Array<ChinchiroGame>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  participants: Array<ChinchiroRoomParticipant>;
  roomMasters: Array<ChinchiroRoomMaster>;
  settings: ChinchiroRoomSettings;
  status: ChinchiroRoomStatus;
};

export type ChinchiroRoomMaster = {
  __typename?: 'ChinchiroRoomMaster';
  id: Scalars['ID']['output'];
  player: Player;
};

export type ChinchiroRoomParticipant = {
  __typename?: 'ChinchiroRoomParticipant';
  id: Scalars['ID']['output'];
  isGone: Scalars['Boolean']['output'];
  name: Scalars['String']['output'];
  player: Player;
};

export type ChinchiroRoomSettings = {
  __typename?: 'ChinchiroRoomSettings';
  dummy: Scalars['String']['output'];
};

export enum ChinchiroRoomStatus {
  Finished = 'Finished',
  Opened = 'Opened'
}

export type ChinchiroRoomsQuery = {
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  name?: InputMaybe<Scalars['String']['input']>;
  paging?: InputMaybe<PageableQuery>;
  statuses?: InputMaybe<Array<ChinchiroRoomStatus>>;
};

export type DeleteChinchiroRoomMaster = {
  masterId: Scalars['ID']['input'];
  roomID: Scalars['ID']['input'];
};

export type DeleteChinchiroRoomMasterPayload = {
  __typename?: 'DeleteChinchiroRoomMasterPayload';
  ok: Scalars['Boolean']['output'];
};

export type DeleteChinchiroRoomParticipant = {
  participantId: Scalars['ID']['input'];
  roomId: Scalars['ID']['input'];
};

export type DeleteChinchiroRoomParticipantPayload = {
  __typename?: 'DeleteChinchiroRoomParticipantPayload';
  ok: Scalars['Boolean']['output'];
};

export type LeaveChinchiroRoom = {
  roomId: Scalars['ID']['input'];
};

export type LeaveChinchiroRoomPayload = {
  __typename?: 'LeaveChinchiroRoomPayload';
  ok: Scalars['Boolean']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  betChinchiroGameTurnParticipant: BetChinchiroGameTurnParticipantPayload;
  deleteChinchiroRoomMaster: DeleteChinchiroRoomMasterPayload;
  deleteChinchiroRoomParticipant: DeleteChinchiroRoomParticipantPayload;
  leaveChinchiroRoom: LeaveChinchiroRoomPayload;
  registerChinchiroGame: RegisterChinchiroGamePayload;
  registerChinchiroRoom: RegisterChinchiroRoomPayload;
  registerChinchiroRoomMaster: RegisterChinchiroRoomMasterPayload;
  registerChinchiroRoomParticipant: RegisterChinchiroRoomParticipantPayload;
  rollChinchiroGameTurnParticipant: RollChinchiroGameTurnParticipantPayload;
  updateChinchiroGameTurnStatus: UpdateChinchiroGameTurnStatusPayload;
  updateChinchiroRoomParticipant: UpdateChinchiroRoomParticipantPayload;
  updateChinchiroRoomSettings: UpdateChinchiroRoomSettingsPayload;
  updateChinchiroRoomStatus: UpdateChinchiroRoomStatusPayload;
  updatePlayerProfile: UpdatePlayerProfilePayload;
};


export type MutationBetChinchiroGameTurnParticipantArgs = {
  input: BetChinchiroGameTurnParticipant;
};


export type MutationDeleteChinchiroRoomMasterArgs = {
  input: DeleteChinchiroRoomMaster;
};


export type MutationDeleteChinchiroRoomParticipantArgs = {
  input: DeleteChinchiroRoomParticipant;
};


export type MutationLeaveChinchiroRoomArgs = {
  input: LeaveChinchiroRoom;
};


export type MutationRegisterChinchiroGameArgs = {
  input: NewChinchiroGame;
};


export type MutationRegisterChinchiroRoomArgs = {
  input: NewChinchiroRoom;
};


export type MutationRegisterChinchiroRoomMasterArgs = {
  input: NewChinchiroRoomMaster;
};


export type MutationRegisterChinchiroRoomParticipantArgs = {
  input: NewChinchiroRoomParticipant;
};


export type MutationRollChinchiroGameTurnParticipantArgs = {
  input: RollChinchiroGameTurnParticipant;
};


export type MutationUpdateChinchiroGameTurnStatusArgs = {
  input: UpdateChinchiroGameTurnStatus;
};


export type MutationUpdateChinchiroRoomParticipantArgs = {
  input: UpdateChinchiroRoomParticipant;
};


export type MutationUpdateChinchiroRoomSettingsArgs = {
  input: UpdateChinchiroRoomSettings;
};


export type MutationUpdateChinchiroRoomStatusArgs = {
  input: UpdateChinchiroRoomStatus;
};


export type MutationUpdatePlayerProfileArgs = {
  input: UpdatePlayerProfile;
};

export type NewChinchiroGame = {
  roomId: Scalars['ID']['input'];
};

export type NewChinchiroRoom = {
  name: Scalars['String']['input'];
};

export type NewChinchiroRoomMaster = {
  playerId: Scalars['ID']['input'];
  roomId: Scalars['ID']['input'];
};

export type NewChinchiroRoomParticipant = {
  name: Scalars['String']['input'];
  password?: InputMaybe<Scalars['String']['input']>;
  roomId: Scalars['ID']['input'];
};

export type Pageable = {
  allPageCount: Scalars['Int']['output'];
  currentPageNumber?: Maybe<Scalars['Int']['output']>;
  hasNextPage: Scalars['Boolean']['output'];
  hasPrePage: Scalars['Boolean']['output'];
  isDesc: Scalars['Boolean']['output'];
};

export type PageableQuery = {
  isDesc: Scalars['Boolean']['input'];
  isLatest: Scalars['Boolean']['input'];
  pageNumber: Scalars['Int']['input'];
  pageSize: Scalars['Int']['input'];
};

export type Player = {
  __typename?: 'Player';
  authorityCodes: Array<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type PlayersQuery = {
  ids?: InputMaybe<Array<Scalars['ID']['input']>>;
  name?: InputMaybe<Scalars['String']['input']>;
  paging?: InputMaybe<PageableQuery>;
};

export type Query = {
  __typename?: 'Query';
  chinchiroGame?: Maybe<ChinchiroGame>;
  chinchiroGameTurn?: Maybe<ChinchiroGameTurn>;
  chinchiroGameTurnParticipantResults: Array<ChinchiroGameTurnParticipantResult>;
  chinchiroGameTurnRolls: Array<ChinchiroGameTurnParticipantRoll>;
  chinchiroGameTurns: Array<ChinchiroGameTurn>;
  chinchiroGames: Array<ChinchiroGame>;
  chinchiroRoom?: Maybe<ChinchiroRoom>;
  chinchiroRooms: Array<SimpleChinchiroRoom>;
  myChinchiroGameParticipant?: Maybe<ChinchiroGameParticipant>;
  myChinchiroRoomParticipant?: Maybe<ChinchiroRoomParticipant>;
  myPlayer?: Maybe<Player>;
  player?: Maybe<Player>;
  players: Array<Player>;
};


export type QueryChinchiroGameArgs = {
  gameId: Scalars['ID']['input'];
};


export type QueryChinchiroGameTurnArgs = {
  turnId: Scalars['ID']['input'];
};


export type QueryChinchiroGameTurnParticipantResultsArgs = {
  query: ChinchiroGameTurnParticipantResultsQuery;
};


export type QueryChinchiroGameTurnRollsArgs = {
  query?: InputMaybe<ChinchiroGameTurnRollsQuery>;
};


export type QueryChinchiroGameTurnsArgs = {
  query: ChinchiroGameTurnsQuery;
};


export type QueryChinchiroGamesArgs = {
  query: ChinchiroGamesQuery;
};


export type QueryChinchiroRoomArgs = {
  roomId: Scalars['ID']['input'];
};


export type QueryChinchiroRoomsArgs = {
  query: ChinchiroRoomsQuery;
};


export type QueryMyChinchiroGameParticipantArgs = {
  gameId: Scalars['ID']['input'];
};


export type QueryMyChinchiroRoomParticipantArgs = {
  roomId: Scalars['ID']['input'];
};


export type QueryPlayerArgs = {
  id: Scalars['ID']['input'];
};


export type QueryPlayersArgs = {
  query: PlayersQuery;
};

export type RegisterChinchiroGamePayload = {
  __typename?: 'RegisterChinchiroGamePayload';
  chinchiroGame: ChinchiroGame;
};

export type RegisterChinchiroRoomMasterPayload = {
  __typename?: 'RegisterChinchiroRoomMasterPayload';
  chinchiroRoomMaster: ChinchiroRoomMaster;
};

export type RegisterChinchiroRoomParticipantPayload = {
  __typename?: 'RegisterChinchiroRoomParticipantPayload';
  chinchiroRoomParticipant: ChinchiroRoomParticipant;
};

export type RegisterChinchiroRoomPayload = {
  __typename?: 'RegisterChinchiroRoomPayload';
  chinchiroRoom: ChinchiroRoom;
};

export type RollChinchiroGameTurnParticipant = {
  turnId: Scalars['ID']['input'];
};

export type RollChinchiroGameTurnParticipantPayload = {
  __typename?: 'RollChinchiroGameTurnParticipantPayload';
  ok: Scalars['Boolean']['output'];
};

export type SimpleChinchiroRoom = {
  __typename?: 'SimpleChinchiroRoom';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  participantCounts: Scalars['Int']['output'];
  settings: ChinchiroRoomSettings;
  status: ChinchiroRoomStatus;
};

export type UpdateChinchiroGameTurnStatus = {
  status: ChinchiroGameTurnStatus;
  turnId: Scalars['ID']['input'];
};

export type UpdateChinchiroGameTurnStatusPayload = {
  __typename?: 'UpdateChinchiroGameTurnStatusPayload';
  ok: Scalars['Boolean']['output'];
};

export type UpdateChinchiroRoomParticipant = {
  name: Scalars['String']['input'];
  participantId: Scalars['ID']['input'];
  roomId: Scalars['ID']['input'];
};

export type UpdateChinchiroRoomParticipantPayload = {
  __typename?: 'UpdateChinchiroRoomParticipantPayload';
  ok: Scalars['Boolean']['output'];
};

export type UpdateChinchiroRoomSettings = {
  dummy: Scalars['String']['input'];
  roomId: Scalars['ID']['input'];
};

export type UpdateChinchiroRoomSettingsPayload = {
  __typename?: 'UpdateChinchiroRoomSettingsPayload';
  ok: Scalars['Boolean']['output'];
};

export type UpdateChinchiroRoomStatus = {
  roomId: Scalars['ID']['input'];
  status: ChinchiroRoomStatus;
};

export type UpdateChinchiroRoomStatusPayload = {
  __typename?: 'UpdateChinchiroRoomStatusPayload';
  ok: Scalars['Boolean']['output'];
};

export type UpdatePlayerProfile = {
  name: Scalars['String']['input'];
};

export type UpdatePlayerProfilePayload = {
  __typename?: 'UpdatePlayerProfilePayload';
  ok: Scalars['Boolean']['output'];
};

export type MleaveChinchiroRoomMutationVariables = Exact<{
  input: LeaveChinchiroRoom;
}>;


export type MleaveChinchiroRoomMutation = { __typename?: 'Mutation', leaveChinchiroRoom: { __typename?: 'LeaveChinchiroRoomPayload', ok: boolean } };

export type ParticipateChinchiroRoomMutationVariables = Exact<{
  input: NewChinchiroRoomParticipant;
}>;


export type ParticipateChinchiroRoomMutation = { __typename?: 'Mutation', registerChinchiroRoomParticipant: { __typename?: 'RegisterChinchiroRoomParticipantPayload', chinchiroRoomParticipant: { __typename?: 'ChinchiroRoomParticipant', id: string } } };

export type RegisterChinchiroGameMutationVariables = Exact<{
  input: NewChinchiroGame;
}>;


export type RegisterChinchiroGameMutation = { __typename?: 'Mutation', registerChinchiroGame: { __typename?: 'RegisterChinchiroGamePayload', chinchiroGame: { __typename?: 'ChinchiroGame', id: string } } };

export type RegisterChinchiroRoomMutationVariables = Exact<{
  input: NewChinchiroRoom;
}>;


export type RegisterChinchiroRoomMutation = { __typename?: 'Mutation', registerChinchiroRoom: { __typename?: 'RegisterChinchiroRoomPayload', chinchiroRoom: { __typename?: 'ChinchiroRoom', id: string } } };

export type UpdatePlayerProfileMutationVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type UpdatePlayerProfileMutation = { __typename?: 'Mutation', updatePlayerProfile: { __typename?: 'UpdatePlayerProfilePayload', ok: boolean } };

export type QChinchiroGamesQueryVariables = Exact<{
  query: ChinchiroGamesQuery;
}>;


export type QChinchiroGamesQuery = { __typename?: 'Query', chinchiroGames: Array<{ __typename?: 'ChinchiroGame', id: string, status: ChinchiroGameStatus, participants: Array<{ __typename?: 'ChinchiroGameParticipant', id: string, balance: number, turnOrder: number, roomParticipant: { __typename?: 'ChinchiroRoomParticipant', id: string } }>, turns: Array<{ __typename?: 'ChinchiroGameTurn', id: string, status: ChinchiroGameTurnStatus, turnNumber: number, dealer: { __typename?: 'ChinchiroGameParticipant', id: string, turnOrder: number }, nextRoller?: { __typename?: 'ChinchiroGameParticipant', id: string, turnOrder: number } | null, rolls: Array<{ __typename?: 'ChinchiroGameTurnParticipantRoll', id: string, rollNumber: number, dice1: number, dice2: number, dice3: number, participant: { __typename?: 'ChinchiroGameParticipant', id: string } }>, results: Array<{ __typename?: 'ChinchiroGameTurnParticipantResult', id: string, dice1: number, dice2: number, dice3: number, combination: ChinchiroCombination, winnings: number, participant: { __typename?: 'ChinchiroGameParticipant', id: string } }> }> }> };

export type QChinchiroRoomQueryVariables = Exact<{
  roomId: Scalars['ID']['input'];
}>;


export type QChinchiroRoomQuery = { __typename?: 'Query', chinchiroRoom?: { __typename?: 'ChinchiroRoom', id: string, name: string, status: ChinchiroRoomStatus, roomMasters: Array<{ __typename?: 'ChinchiroRoomMaster', id: string, player: { __typename?: 'Player', id: string, name: string } }>, participants: Array<{ __typename?: 'ChinchiroRoomParticipant', id: string, name: string, player: { __typename?: 'Player', id: string, name: string } }>, games: Array<{ __typename?: 'ChinchiroGame', id: string, status: ChinchiroGameStatus }>, settings: { __typename?: 'ChinchiroRoomSettings', dummy: string } } | null };

export type QChinchiroRoomsQueryVariables = Exact<{
  query: ChinchiroRoomsQuery;
}>;


export type QChinchiroRoomsQuery = { __typename?: 'Query', chinchiroRooms: Array<{ __typename?: 'SimpleChinchiroRoom', id: string, name: string, status: ChinchiroRoomStatus }> };

export type MyChinchiroRoomParticipantQueryVariables = Exact<{
  roomId: Scalars['ID']['input'];
}>;


export type MyChinchiroRoomParticipantQuery = { __typename?: 'Query', myChinchiroRoomParticipant?: { __typename?: 'ChinchiroRoomParticipant', id: string, name: string, player: { __typename?: 'Player', id: string, name: string } } | null };

export type MyPlayerQueryVariables = Exact<{ [key: string]: never; }>;


export type MyPlayerQuery = { __typename?: 'Query', myPlayer?: { __typename?: 'Player', id: string, name: string, authorityCodes: Array<string> } | null };

export type PlayerQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type PlayerQuery = { __typename?: 'Query', player?: { __typename?: 'Player', id: string, name: string } | null };

export type QPlayersQueryVariables = Exact<{
  query: PlayersQuery;
}>;


export type QPlayersQuery = { __typename?: 'Query', players: Array<{ __typename?: 'Player', id: string, name: string }> };


export const MleaveChinchiroRoomDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"MleaveChinchiroRoom"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"LeaveChinchiroRoom"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"leaveChinchiroRoom"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ok"}}]}}]}}]} as unknown as DocumentNode<MleaveChinchiroRoomMutation, MleaveChinchiroRoomMutationVariables>;
export const ParticipateChinchiroRoomDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"participateChinchiroRoom"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewChinchiroRoomParticipant"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"registerChinchiroRoomParticipant"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroRoomParticipant"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<ParticipateChinchiroRoomMutation, ParticipateChinchiroRoomMutationVariables>;
export const RegisterChinchiroGameDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"registerChinchiroGame"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewChinchiroGame"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"registerChinchiroGame"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroGame"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<RegisterChinchiroGameMutation, RegisterChinchiroGameMutationVariables>;
export const RegisterChinchiroRoomDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"registerChinchiroRoom"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewChinchiroRoom"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"registerChinchiroRoom"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroRoom"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<RegisterChinchiroRoomMutation, RegisterChinchiroRoomMutationVariables>;
export const UpdatePlayerProfileDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdatePlayerProfile"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updatePlayerProfile"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ok"}}]}}]}}]} as unknown as DocumentNode<UpdatePlayerProfileMutation, UpdatePlayerProfileMutationVariables>;
export const QChinchiroGamesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"QChinchiroGames"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"query"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChinchiroGamesQuery"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroGames"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"query"},"value":{"kind":"Variable","name":{"kind":"Name","value":"query"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"participants"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"roomParticipant"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"balance"}},{"kind":"Field","name":{"kind":"Name","value":"turnOrder"}}]}},{"kind":"Field","name":{"kind":"Name","value":"turns"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"dealer"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"turnOrder"}}]}},{"kind":"Field","name":{"kind":"Name","value":"nextRoller"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"turnOrder"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"turnNumber"}},{"kind":"Field","name":{"kind":"Name","value":"rolls"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"participant"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rollNumber"}},{"kind":"Field","name":{"kind":"Name","value":"dice1"}},{"kind":"Field","name":{"kind":"Name","value":"dice2"}},{"kind":"Field","name":{"kind":"Name","value":"dice3"}}]}},{"kind":"Field","name":{"kind":"Name","value":"results"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"participant"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"dice1"}},{"kind":"Field","name":{"kind":"Name","value":"dice2"}},{"kind":"Field","name":{"kind":"Name","value":"dice3"}},{"kind":"Field","name":{"kind":"Name","value":"combination"}},{"kind":"Field","name":{"kind":"Name","value":"winnings"}}]}}]}}]}}]}}]} as unknown as DocumentNode<QChinchiroGamesQuery, QChinchiroGamesQueryVariables>;
export const QChinchiroRoomDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"QChinchiroRoom"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"roomId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroRoom"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"roomId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"roomId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"roomMasters"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"player"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"participants"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"player"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"games"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"settings"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"dummy"}}]}}]}}]}}]} as unknown as DocumentNode<QChinchiroRoomQuery, QChinchiroRoomQueryVariables>;
export const QChinchiroRoomsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"QChinchiroRooms"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"query"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChinchiroRoomsQuery"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chinchiroRooms"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"query"},"value":{"kind":"Variable","name":{"kind":"Name","value":"query"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<QChinchiroRoomsQuery, QChinchiroRoomsQueryVariables>;
export const MyChinchiroRoomParticipantDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"MyChinchiroRoomParticipant"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"roomId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"myChinchiroRoomParticipant"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"roomId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"roomId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"player"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<MyChinchiroRoomParticipantQuery, MyChinchiroRoomParticipantQueryVariables>;
export const MyPlayerDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"MyPlayer"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"myPlayer"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"authorityCodes"}}]}}]}}]} as unknown as DocumentNode<MyPlayerQuery, MyPlayerQueryVariables>;
export const PlayerDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Player"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"player"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<PlayerQuery, PlayerQueryVariables>;
export const QPlayersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"QPlayers"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"query"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"PlayersQuery"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"players"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"query"},"value":{"kind":"Variable","name":{"kind":"Name","value":"query"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<QPlayersQuery, QPlayersQueryVariables>;