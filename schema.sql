-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Feb 25, 2020 at 09:00 PM
-- Server version: 5.7.26
-- PHP Version: 7.3.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `rssDB`
--
CREATE DATABASE rssDB;
-- --------------------------------------------------------

--
-- Table structure for table `feed`
--

USE rssDB;

CREATE TABLE `feed` (
  `id` int(255) NOT NULL,
  `title` varchar(500) NOT NULL,
  `description` text NOT NULL,
  `link` varchar(500) NOT NULL,
  `published_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `feed`
--
ALTER TABLE `feed`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `feed` ADD FULLTEXT KEY `FULL_TEXT` (`title`,`description`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `feed`
--
ALTER TABLE `feed`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT;
