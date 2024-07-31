/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "mutation MleaveChinchiroRoom($input: LeaveChinchiroRoom!) {\n  leaveChinchiroRoom(input: $input) {\n    ok\n  }\n}": types.MleaveChinchiroRoomDocument,
    "mutation participateChinchiroRoom($input: NewChinchiroRoomParticipant!) {\n  registerChinchiroRoomParticipant(input: $input) {\n    chinchiroRoomParticipant {\n      id\n    }\n  }\n}": types.ParticipateChinchiroRoomDocument,
    "mutation registerChinchiroGame($input: NewChinchiroGame!) {\n  registerChinchiroGame(input: $input) {\n    chinchiroGame {\n      id\n    }\n  }\n}": types.RegisterChinchiroGameDocument,
    "mutation registerChinchiroRoom($input: NewChinchiroRoom!) {\n  registerChinchiroRoom(input: $input) {\n    chinchiroRoom {\n      id\n    }\n  }\n}": types.RegisterChinchiroRoomDocument,
    "mutation UpdatePlayerProfile($name: String!) {\n  updatePlayerProfile(input: {name: $name}) {\n    ok\n  }\n}": types.UpdatePlayerProfileDocument,
    "query QChinchiroGames($query: ChinchiroGamesQuery!) {\n  chinchiroGames(query: $query) {\n    id\n    status\n    participants {\n      id\n      roomParticipant {\n        id\n      }\n      balance\n      turnOrder\n    }\n    turns {\n      id\n      dealer {\n        id\n        turnOrder\n      }\n      nextRoller {\n        id\n        turnOrder\n      }\n      status\n      turnNumber\n      rolls {\n        id\n        participant {\n          id\n        }\n        rollNumber\n        dice1\n        dice2\n        dice3\n      }\n      results {\n        id\n        participant {\n          id\n        }\n        dice1\n        dice2\n        dice3\n        combination\n        winnings\n      }\n    }\n  }\n}": types.QChinchiroGamesDocument,
    "query QChinchiroRoom($roomId: ID!) {\n  chinchiroRoom(roomId: $roomId) {\n    id\n    name\n    status\n    roomMasters {\n      id\n      player {\n        id\n        name\n      }\n    }\n    participants {\n      id\n      name\n      player {\n        id\n        name\n      }\n    }\n    games {\n      id\n      status\n    }\n    settings {\n      dummy\n    }\n  }\n}": types.QChinchiroRoomDocument,
    "query QChinchiroRooms($query: ChinchiroRoomsQuery!) {\n  chinchiroRooms(query: $query) {\n    id\n    name\n    status\n  }\n}": types.QChinchiroRoomsDocument,
    "query MyChinchiroRoomParticipant($roomId: ID!) {\n  myChinchiroRoomParticipant(roomId: $roomId) {\n    id\n    name\n    player {\n      id\n      name\n    }\n  }\n}": types.MyChinchiroRoomParticipantDocument,
    "query MyPlayer {\n  myPlayer {\n    id\n    name\n    authorityCodes\n  }\n}": types.MyPlayerDocument,
    "query Player($id: ID!) {\n  player(id: $id) {\n    id\n    name\n  }\n}": types.PlayerDocument,
    "query QPlayers($query: PlayersQuery!) {\n  players(query: $query) {\n    id\n    name\n  }\n}": types.QPlayersDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation MleaveChinchiroRoom($input: LeaveChinchiroRoom!) {\n  leaveChinchiroRoom(input: $input) {\n    ok\n  }\n}"): (typeof documents)["mutation MleaveChinchiroRoom($input: LeaveChinchiroRoom!) {\n  leaveChinchiroRoom(input: $input) {\n    ok\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation participateChinchiroRoom($input: NewChinchiroRoomParticipant!) {\n  registerChinchiroRoomParticipant(input: $input) {\n    chinchiroRoomParticipant {\n      id\n    }\n  }\n}"): (typeof documents)["mutation participateChinchiroRoom($input: NewChinchiroRoomParticipant!) {\n  registerChinchiroRoomParticipant(input: $input) {\n    chinchiroRoomParticipant {\n      id\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation registerChinchiroGame($input: NewChinchiroGame!) {\n  registerChinchiroGame(input: $input) {\n    chinchiroGame {\n      id\n    }\n  }\n}"): (typeof documents)["mutation registerChinchiroGame($input: NewChinchiroGame!) {\n  registerChinchiroGame(input: $input) {\n    chinchiroGame {\n      id\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation registerChinchiroRoom($input: NewChinchiroRoom!) {\n  registerChinchiroRoom(input: $input) {\n    chinchiroRoom {\n      id\n    }\n  }\n}"): (typeof documents)["mutation registerChinchiroRoom($input: NewChinchiroRoom!) {\n  registerChinchiroRoom(input: $input) {\n    chinchiroRoom {\n      id\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation UpdatePlayerProfile($name: String!) {\n  updatePlayerProfile(input: {name: $name}) {\n    ok\n  }\n}"): (typeof documents)["mutation UpdatePlayerProfile($name: String!) {\n  updatePlayerProfile(input: {name: $name}) {\n    ok\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query QChinchiroGames($query: ChinchiroGamesQuery!) {\n  chinchiroGames(query: $query) {\n    id\n    status\n    participants {\n      id\n      roomParticipant {\n        id\n      }\n      balance\n      turnOrder\n    }\n    turns {\n      id\n      dealer {\n        id\n        turnOrder\n      }\n      nextRoller {\n        id\n        turnOrder\n      }\n      status\n      turnNumber\n      rolls {\n        id\n        participant {\n          id\n        }\n        rollNumber\n        dice1\n        dice2\n        dice3\n      }\n      results {\n        id\n        participant {\n          id\n        }\n        dice1\n        dice2\n        dice3\n        combination\n        winnings\n      }\n    }\n  }\n}"): (typeof documents)["query QChinchiroGames($query: ChinchiroGamesQuery!) {\n  chinchiroGames(query: $query) {\n    id\n    status\n    participants {\n      id\n      roomParticipant {\n        id\n      }\n      balance\n      turnOrder\n    }\n    turns {\n      id\n      dealer {\n        id\n        turnOrder\n      }\n      nextRoller {\n        id\n        turnOrder\n      }\n      status\n      turnNumber\n      rolls {\n        id\n        participant {\n          id\n        }\n        rollNumber\n        dice1\n        dice2\n        dice3\n      }\n      results {\n        id\n        participant {\n          id\n        }\n        dice1\n        dice2\n        dice3\n        combination\n        winnings\n      }\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query QChinchiroRoom($roomId: ID!) {\n  chinchiroRoom(roomId: $roomId) {\n    id\n    name\n    status\n    roomMasters {\n      id\n      player {\n        id\n        name\n      }\n    }\n    participants {\n      id\n      name\n      player {\n        id\n        name\n      }\n    }\n    games {\n      id\n      status\n    }\n    settings {\n      dummy\n    }\n  }\n}"): (typeof documents)["query QChinchiroRoom($roomId: ID!) {\n  chinchiroRoom(roomId: $roomId) {\n    id\n    name\n    status\n    roomMasters {\n      id\n      player {\n        id\n        name\n      }\n    }\n    participants {\n      id\n      name\n      player {\n        id\n        name\n      }\n    }\n    games {\n      id\n      status\n    }\n    settings {\n      dummy\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query QChinchiroRooms($query: ChinchiroRoomsQuery!) {\n  chinchiroRooms(query: $query) {\n    id\n    name\n    status\n  }\n}"): (typeof documents)["query QChinchiroRooms($query: ChinchiroRoomsQuery!) {\n  chinchiroRooms(query: $query) {\n    id\n    name\n    status\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query MyChinchiroRoomParticipant($roomId: ID!) {\n  myChinchiroRoomParticipant(roomId: $roomId) {\n    id\n    name\n    player {\n      id\n      name\n    }\n  }\n}"): (typeof documents)["query MyChinchiroRoomParticipant($roomId: ID!) {\n  myChinchiroRoomParticipant(roomId: $roomId) {\n    id\n    name\n    player {\n      id\n      name\n    }\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query MyPlayer {\n  myPlayer {\n    id\n    name\n    authorityCodes\n  }\n}"): (typeof documents)["query MyPlayer {\n  myPlayer {\n    id\n    name\n    authorityCodes\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query Player($id: ID!) {\n  player(id: $id) {\n    id\n    name\n  }\n}"): (typeof documents)["query Player($id: ID!) {\n  player(id: $id) {\n    id\n    name\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query QPlayers($query: PlayersQuery!) {\n  players(query: $query) {\n    id\n    name\n  }\n}"): (typeof documents)["query QPlayers($query: PlayersQuery!) {\n  players(query: $query) {\n    id\n    name\n  }\n}"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;