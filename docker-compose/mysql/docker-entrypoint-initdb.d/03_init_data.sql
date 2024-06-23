use gamesdb;

insert into players (player_name, created_at, updated_at) values ('player1', now(), now());
insert into player_accounts (player_id, user_name, created_at, updated_at) values (1, 'user1', now(), now());
insert into player_authorities (player_id, authority_code, created_at, updated_at) values (1, 'ROLE_USER', now(), now());

insert into players (player_name, created_at, updated_at) values ('player2', now(), now());
insert into player_accounts (player_id, user_name, created_at, updated_at) values (2, 'user2', now(), now());
insert into player_authorities (player_id, authority_code, created_at, updated_at) values (2, 'ROLE_USER', now(), now());

-- 部屋1
insert into chinchiro_rooms (room_name, room_status_code, created_at, updated_at) values ('room1', 'Opened', now(), now());
insert into chinchiro_room_master_players (room_id, player_id, created_at, updated_at) values (1, 1, now(), now());

insert into chinchiro_room_participants (room_id, player_id, participant_name, is_gone, created_at, updated_at) values (1, 1, 'participant_name1', false, now(), now());
insert into chinchiro_room_participants (room_id, player_id, participant_name, is_gone, created_at, updated_at) values (1, 2, 'participant_name2', false, now(), now());

insert into chinchiro_room_settings (room_id, room_setting_key, room_setting_value, created_at, updated_at) values (1, 'Password', 'password', now(), now());

-- ゲーム1
insert into chinchiro_games (room_id, game_status_code, created_at, updated_at) values (1, 'Progress', now(), now());

insert into chinchiro_game_participants (game_id, room_participant_id, balance, turn_order, created_at, updated_at) values (1, 1, 100000, 1, now(), now());
insert into chinchiro_game_participants (game_id, room_participant_id, balance, turn_order, created_at, updated_at) values (1, 2, 100000, 2, now(), now());

-- ターン1（完了）
insert into chinchiro_game_turns (game_id, dealer_participant_id, next_roller_participant_id, turn_status_code, turn_number, created_at, updated_at) values (1, 1, null, 'Finished', 1, now(), now());
-- 一人目のロール
insert into chinchiro_game_turn_participant_rolls (game_id, turn_id, participant_id, roll_number, dice1, dice2, dice3, created_at, updated_at) values (1, 1, 1, 1, 2, 3, 4, now(), now());
insert into chinchiro_game_turn_participant_rolls (game_id, turn_id, participant_id, roll_number, dice1, dice2, dice3, created_at, updated_at) values (1, 1, 1, 2, 2, 2, 4, now(), now());
-- 二人目のロール
insert into chinchiro_game_turn_participant_rolls (game_id, turn_id, participant_id, roll_number, dice1, dice2, dice3, created_at, updated_at) values (1, 1, 2, 1, 1, 1, 1, now(), now());
-- 結果
insert into chinchiro_game_turn_participant_results (game_id, turn_id, participant_id, bet_amount, dice1, dice2, dice3, combination_code, winnings, created_at, updated_at) 
    values (1, 1, 1, null, 2, 2, 4, 'Yonnome', -20000, now(), now());
insert into chinchiro_game_turn_participant_results (game_id, turn_id, participant_id, bet_amount, dice1, dice2, dice3, combination_code, winnings, created_at, updated_at) 
    values (1, 1, 2, null, 1, 1, 1, 'Pinzoro', 20000, now(), now());

-- ターン2（ロール中）
insert into chinchiro_game_turns (game_id, dealer_participant_id, next_roller_participant_id, turn_status_code, turn_number, created_at, updated_at) values (1, 2, 1, 'Rolling', 2, now(), now());
-- 結果
insert into chinchiro_game_turn_participant_results (game_id, turn_id, participant_id, bet_amount, dice1, dice2, dice3, combination_code, winnings, created_at, updated_at) 
    values (1, 2, 1, 10000, null, null, null, null, null, now(), now());
insert into chinchiro_game_turn_participant_results (game_id, turn_id, participant_id, bet_amount, dice1, dice2, dice3, combination_code, winnings, created_at, updated_at) 
    values (1, 2, 2, null, null, null, null, null, null, now(), now());

-- ゲーム2
insert into chinchiro_games (room_id, game_status_code, created_at, updated_at) values (1, 'Progress', now(), now());

insert into chinchiro_game_participants (game_id, room_participant_id, balance, turn_order, created_at, updated_at) values (2, 1, 100000, 1, now(), now());
insert into chinchiro_game_participants (game_id, room_participant_id, balance, turn_order, created_at, updated_at) values (2, 2, 100000, 2, now(), now());

-- ターン1（ベット中）
insert into chinchiro_game_turns (game_id, dealer_participant_id, next_roller_participant_id, turn_status_code, turn_number, created_at, updated_at) values (2, 1, null, 'Betting', 1, now(), now());

-- ゲーム3（参加者募集中）
insert into chinchiro_games (room_id, game_status_code, created_at, updated_at) values (1, 'Progress', now(), now());
