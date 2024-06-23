use gamesdb;


create table players (
    id   int unsigned not null auto_increment comment 'ID',
    player_name varchar(255) not null comment 'プレイヤー名',
    created_at  datetime     not null comment '作成日時',
    updated_at  datetime     not null comment '更新日時',
    primary key (id)
);

create table player_accounts (
    id         int unsigned not null auto_increment comment 'ID',
    player_id  int unsigned not null,
    user_name  varchar(255) not null unique comment 'ユーザー名',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id)
);

alter table player_accounts
    add constraint fk_player_accounts_players foreign key (player_id)
    references players (id)
    on update restrict
    on delete restrict
;

create table player_authorities (
    id                  int unsigned not null auto_increment comment 'ID',
    player_id           int unsigned not null,
    authority_code      varchar(255) not null comment '権限コード',
    created_at datetime not null comment '作成日時',
    updated_at datetime not null comment '更新日時',
    primary key (id),
    constraint uq_player_authority unique (player_id, authority_code)
);

alter table player_authorities
    add constraint fk_player_authorities_players foreign key (player_id)
    references players (id)
    on update restrict
    on delete restrict
;

-- chinchiro

create table chinchiro_rooms (
    id               int unsigned not null auto_increment comment 'ID',
    room_name        varchar(255) not null comment 'ルーム名',
    room_status_code varchar(255) not null comment 'ルームステータスコード',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id)
);

create table chinchiro_room_master_players (
    id          int unsigned not null auto_increment comment 'ID',
    room_id     int unsigned not null comment 'ルームID',
    player_id   int unsigned not null comment 'プレイヤーID',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (room_id, player_id)
);

alter table chinchiro_room_master_players
    add constraint fk_chinchiro_room_master_players_rooms foreign key (room_id)
    references chinchiro_rooms (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_room_master_players
    add constraint fk_chinchiro_room_master_players_players foreign key (player_id)
    references players (id)
    on update restrict
    on delete restrict
;

create table chinchiro_room_participants (
    id                    int unsigned not null auto_increment comment 'ID',
    room_id               int unsigned not null comment 'ルームID',
    player_id             int unsigned not null comment 'プレイヤーID',
    participant_name      varchar(255) not null comment 'ゲーム参加者名',
    is_gone               boolean      not null comment 'ゲームから退出したか',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id)
);

alter table chinchiro_room_participants
    add constraint fk_chinchiro_room_participants_chinchiro_rooms foreign key (room_id)
    references chinchiro_rooms (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_room_participants
    add constraint fk_chinchiro_room_participants_players foreign key (player_id)
    references players (id)
    on update restrict
    on delete restrict
;

create table chinchiro_room_settings (
    id                 int unsigned not null auto_increment comment 'ID',
    room_id            int unsigned not null comment 'ルームID',
    room_setting_key   varchar(255) not null comment '設定キー',
    room_setting_value text         not null comment '設定値',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (room_id, room_setting_key)
);

alter table chinchiro_room_settings
    add constraint fk_chinchiro_room_settings_chinchiro_rooms foreign key (room_id)
    references chinchiro_rooms (id)
    on update restrict
    on delete restrict
;

create table chinchiro_games (
    id               int unsigned not null auto_increment comment 'ID',
    room_id          int unsigned not null comment 'ルームID',
    game_status_code varchar(255) not null comment 'ゲームステータスコード',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id)
);

alter table chinchiro_games
    add constraint fk_chinchiro_games_chinchiro_rooms foreign key (room_id)
    references chinchiro_rooms (id)
    on update restrict
    on delete restrict
;

create table chinchiro_game_participants (
    id                   int unsigned not null auto_increment comment 'ID',
    game_id              int unsigned not null comment 'ゲームID',
    room_participant_id  int unsigned not null comment '部屋参加者ID',
    balance              int          not null comment '残高',
    turn_order           int unsigned not null comment '順番',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (game_id, room_participant_id),
    unique (game_id, turn_order)
);

alter table chinchiro_game_participants
    add constraint fk_chinchiro_game_participants_chinchiro_games foreign key (game_id)
    references chinchiro_games (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_participants
    add constraint fk_chinchiro_game_participants_room_participants foreign key (room_participant_id)
    references chinchiro_room_participants (id)
    on update restrict
    on delete restrict
;

create table chinchiro_game_turns (
    id                         int unsigned not null auto_increment comment 'ID',
    game_id                    int unsigned not null comment 'ゲームID',
    dealer_participant_id      int unsigned not null comment 'ディーラーID',
    next_roller_participant_id int unsigned          comment '次のローラーID',
    turn_status_code           varchar(255) not null comment 'ターンステータスコード',
    turn_number                int unsigned not null comment 'ターン番号',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (game_id, turn_number)
);

alter table chinchiro_game_turns
    add constraint fk_chinchiro_game_turns_chinchiro_games foreign key (game_id)
    references chinchiro_games (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_turns
    add constraint fk_chinchiro_game_turns_chinchiro_game_participants foreign key (dealer_participant_id)
    references chinchiro_game_participants (id)
    on update restrict
    on delete restrict
;

create table chinchiro_game_turn_participant_rolls (
    id                    int unsigned not null auto_increment comment 'ID',
    game_id               int unsigned not null comment 'ゲームID',
    turn_id               int unsigned not null comment 'ターンID',
    participant_id        int unsigned not null comment '参加者ID',
    roll_number           int unsigned not null comment 'ロール回数',
    dice1                 int unsigned not null comment 'ダイス1',
    dice2                 int unsigned not null comment 'ダイス2',
    dice3                 int unsigned not null comment 'ダイス3',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (game_id, turn_id, participant_id, roll_number)
);

alter table chinchiro_game_turn_participant_rolls
    add constraint fk_chinchiro_game_turn_participant_rolls_games foreign key (game_id)
    references chinchiro_games (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_turn_participant_rolls
    add constraint fk_chinchiro_game_turn_participant_rolls_turns foreign key (turn_id)
    references chinchiro_game_turns (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_turn_participant_rolls
    add constraint fk_chinchiro_game_turn_participant_rolls_participants foreign key (participant_id)
    references chinchiro_game_participants (id)
    on update restrict
    on delete restrict
;

create table chinchiro_game_turn_participant_results (
    id                    int unsigned not null auto_increment comment 'ID',
    game_id               int unsigned not null comment 'ゲームID',
    turn_id               int unsigned not null comment 'ターンID',
    participant_id        int unsigned not null comment '参加者ID',
    bet_amount            int unsigned          comment '賭け金',
    dice1                 int unsigned          comment 'ダイス1',
    dice2                 int unsigned          comment 'ダイス2',
    dice3                 int unsigned          comment 'ダイス3',
    combination_code      varchar(255)          comment '役コード',
    winnings              int                   comment '獲得金額',
    created_at datetime     not null comment '作成日時',
    updated_at datetime     not null comment '更新日時',
    primary key (id),
    unique (game_id, turn_id, participant_id)
);

alter table chinchiro_game_turn_participant_results
    add constraint fk_chinchiro_game_turn_participant_results_games foreign key (game_id)
    references chinchiro_games (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_turn_participant_results
    add constraint fk_chinchiro_game_turn_participant_results_turns foreign key (turn_id)
    references chinchiro_game_turns (id)
    on update restrict
    on delete restrict
;

alter table chinchiro_game_turn_participant_results
    add constraint fk_chinchiro_game_turn_participant_results_participants foreign key (participant_id)
    references chinchiro_game_participants (id)
    on update restrict
    on delete restrict
;
