create table if not exists articles
(
    article_id integer unsigned auto_increment primary key,
    title      varchar(100) not null,
    contents   text         not null,
    username   varchar(100) not null,
    nice       integer      not null,
    created_at datetime
);

create table if not exists comments
(
    comment_id integer unsigned auto_increment primary key,
    article_id integer unsigned not null,
    message    text             not null,
    created_at datetime,
    foreign key (article_id) references articles (article_id)
);


INSERT INTO articles (article_id, title, contents, username, nice, created_at)
VALUES (1, 'firstPost', 'This is my first blog', 'saki', 2, NOW());
INSERT INTO articles (article_id, title, contents, username, nice, created_at)
VALUES (2, '2nd', 'Second blog post', 'saki', 4, NOW());

INSERT INTO comments (article_id, message, created_at)
VALUES (1, '1st comment yeah', NOW());
INSERT INTO comments (article_id, message, created_at)
VALUES (1, 'welcome', NOW());
