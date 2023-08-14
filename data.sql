--
-- Create a database using `MYSQL_DATABASE` 
--
CREATE DATABASE IF NOT EXISTS `elections`;
USE `elections`;

CREATE TABLE BureauxDeVote (
    Arrondissement int,
    Y2012 int,
    Y2014 int,
    Y2017 int,
    Y2020 int,
    Y2021 int,
    Y2022 int,
);