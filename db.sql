-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 17-04-2023 a las 02:42:54
-- Versión del servidor: 10.4.17-MariaDB
-- Versión de PHP: 8.0.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `parking`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `spots`
--

CREATE TABLE `spots` (
  `id` int(11) NOT NULL,
  `zone` varchar(1) NOT NULL,
  `in_use` tinyint(1) NOT NULL DEFAULT 0,
  `type` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Volcado de datos para la tabla `spots`
--

INSERT INTO `spots` (`id`, `zone`, `in_use`, `type`) VALUES
(1, 'A', 0, 'VIP'),
(2, 'A', 0, 'VIP'),
(3, 'A', 0, 'VIP'),
(4, 'A', 0, 'VIP'),
(5, 'A', 0, 'VIP'),
(6, 'A', 0, 'NORMAL'),
(7, 'A', 0, 'NORMAL'),
(8, 'A', 0, 'NORMAL'),
(9, 'A', 0, 'NORMAL'),
(10, 'A', 0, 'NORMAL'),
(11, 'A', 0, 'NORMAL'),
(12, 'A', 0, 'NORMAL'),
(13, 'A', 0, 'NORMAL'),
(14, 'A', 0, 'NORMAL'),
(15, 'A', 0, 'NORMAL'),
(16, 'A', 0, 'NORMAL'),
(17, 'A', 0, 'NORMAL'),
(18, 'A', 0, 'NORMAL'),
(19, 'A', 0, 'NORMAL'),
(20, 'A', 0, 'NORMAL'),
(21, 'A', 0, 'NORMAL'),
(22, 'A', 0, 'NORMAL'),
(23, 'A', 0, 'NORMAL'),
(24, 'A', 0, 'NORMAL'),
(25, 'A', 0, 'NORMAL'),
(26, 'A', 0, 'NORMAL'),
(27, 'A', 0, 'NORMAL'),
(28, 'A', 0, 'NORMAL'),
(29, 'A', 0, 'NORMAL'),
(30, 'A', 0, 'NORMAL'),
(31, 'B', 0, 'NORMAL'),
(32, 'B', 0, 'NORMAL'),
(33, 'B', 0, 'NORMAL'),
(34, 'B', 0, 'NORMAL'),
(35, 'B', 0, 'NORMAL'),
(36, 'B', 0, 'NORMAL'),
(37, 'B', 0, 'NORMAL'),
(38, 'B', 0, 'NORMAL'),
(39, 'B', 0, 'NORMAL'),
(40, 'B', 0, 'NORMAL'),
(41, 'B', 0, 'NORMAL'),
(42, 'B', 0, 'NORMAL'),
(43, 'B', 0, 'NORMAL'),
(44, 'B', 0, 'NORMAL'),
(45, 'B', 0, 'NORMAL'),
(46, 'B', 0, 'NORMAL'),
(47, 'B', 0, 'NORMAL'),
(48, 'B', 0, 'NORMAL'),
(49, 'B', 0, 'NORMAL'),
(50, 'B', 0, 'NORMAL'),
(51, 'B', 0, 'NORMAL'),
(52, 'B', 0, 'NORMAL'),
(53, 'B', 0, 'NORMAL'),
(54, 'B', 0, 'NORMAL'),
(55, 'B', 0, 'NORMAL'),
(56, 'B', 0, 'NORMAL'),
(57, 'B', 0, 'NORMAL'),
(58, 'B', 0, 'NORMAL'),
(59, 'B', 0, 'NORMAL'),
(60, 'B', 0, 'NORMAL'),
(61, 'C', 0, 'DISCAPACITADO'),
(62, 'C', 0, 'DISCAPACITADO'),
(63, 'C', 0, 'DISCAPACITADO'),
(64, 'C', 0, 'DISCAPACITADO'),
(65, 'C', 0, 'DISCAPACITADO'),
(66, 'C', 0, 'NORMAL'),
(67, 'C', 0, 'NORMAL'),
(68, 'C', 0, 'NORMAL'),
(69, 'C', 0, 'NORMAL'),
(70, 'C', 0, 'NORMAL'),
(71, 'C', 0, 'NORMAL'),
(72, 'C', 0, 'NORMAL'),
(73, 'C', 0, 'NORMAL'),
(74, 'C', 0, 'NORMAL'),
(75, 'C', 0, 'NORMAL'),
(76, 'C', 0, 'NORMAL'),
(77, 'C', 0, 'NORMAL'),
(78, 'C', 0, 'NORMAL'),
(79, 'C', 0, 'NORMAL'),
(80, 'C', 0, 'EMERGENCIA'),
(81, 'C', 0, 'EMERGENCIA'),
(82, 'C', 0, 'EMERGENCIA'),
(83, 'C', 0, 'NORMAL'),
(84, 'C', 0, 'NORMAL'),
(85, 'C', 0, 'NORMAL'),
(86, 'C', 0, 'NORMAL'),
(87, 'C', 0, 'NORMAL'),
(88, 'C', 0, 'NORMAL'),
(89, 'C', 0, 'NORMAL'),
(90, 'C', 0, 'NORMAL'),
(91, 'C', 0, 'NORMAL'),
(92, 'C', 0, 'NORMAL'),
(93, 'C', 0, 'NORMAL'),
(94, 'C', 0, 'NORMAL'),
(95, 'C', 0, 'NORMAL'),
(96, 'C', 0, 'NORMAL'),
(97, 'C', 0, 'NORMAL'),
(98, 'C', 0, 'PROVEEDOR'),
(99, 'C', 0, 'PROVEEDOR'),
(100, 'C', 0, 'PROVEEDOR');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `vehicles`
--

CREATE TABLE `vehicles` (
  `id` int(11) NOT NULL,
  `plate_number` varchar(10) NOT NULL,
  `vehicle_type` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `vehicles_spots`
--

CREATE TABLE `vehicles_spots` (
  `id` int(11) NOT NULL,
  `vehicle_id` int(11) NOT NULL,
  `spot` int(11) NOT NULL,
  `entry_time` datetime NOT NULL DEFAULT current_timestamp(),
  `exit_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `spots`
--
ALTER TABLE `spots`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `vehicles`
--
ALTER TABLE `vehicles`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `vehicles_spots`
--
ALTER TABLE `vehicles_spots`
  ADD PRIMARY KEY (`id`),
  ADD KEY `vehicle_id` (`vehicle_id`,`spot`),
  ADD KEY `spot` (`spot`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `spots`
--
ALTER TABLE `spots`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=101;

--
-- AUTO_INCREMENT de la tabla `vehicles`
--
ALTER TABLE `vehicles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=201;

--
-- AUTO_INCREMENT de la tabla `vehicles_spots`
--
ALTER TABLE `vehicles_spots`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=66;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `vehicles_spots`
--
ALTER TABLE `vehicles_spots`
  ADD CONSTRAINT `vehicles_spots_ibfk_1` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `vehicles_spots_ibfk_2` FOREIGN KEY (`spot`) REFERENCES `spots` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
