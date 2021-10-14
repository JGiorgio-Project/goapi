-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : jeu. 14 oct. 2021 à 10:35
-- Version du serveur :  5.7.31
-- Version de PHP : 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `goapi`
--
CREATE DATABASE IF NOT EXISTS `goapi` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `goapi`;

-- --------------------------------------------------------

--
-- Structure de la table `books`
--

DROP TABLE IF EXISTS `books`;
CREATE TABLE IF NOT EXISTS `books` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Isbn` varchar(255) NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Author` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `books`
--

INSERT INTO `books` (`ID`, `Isbn`, `Title`, `Author`) VALUES
(4, '56887', 'Book Two', 'Steve'),
(3, '432154', 'Book One', 'John'),
(6, '98465', 'Book three', 'Silva'),
(7, '26351', 'Book four', 'Anderson'),
(8, '887546', 'Book five', 'Nael'),
(11, '563142', 'Book seven', 'Amandine'),
(12, '666', 'Book height', 'Helene');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
