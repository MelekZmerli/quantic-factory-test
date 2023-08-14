--
-- Create a database using `MYSQL_DATABASE` 
--
CREATE DATABASE IF NOT EXISTS `elections`;
USE `elections`;

CREATE TABLE BureauxDeVote (
    Arrondissement int,
    y2012 int,
    y2014 int,
    y2017 int,
    y2020 int,
    y2021 int,
    y2022 int,
);