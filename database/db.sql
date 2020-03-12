-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/

CREATE TABLE `category` (
  `id` int(2) NOT NULL,
  `topic` varchar(100) NOT NULL,
  `level` int(1) NOT NULL,
  `level_id` int(2) NOT NULL,
  `status` varchar(1) NOT NULL DEFAULT 'A',
  `position` int(2) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `category`
--

INSERT INTO `category` (`id`, `topic`, `level`, `level_id`, `status`, `position`, `created_at`, `updated_at`) VALUES
(1, 'black tea', 1, 0, 'A', 1, '2019-11-18 11:45:29', '2019-11-23 10:04:34'),
(2, 'green tea', 2, 1, 'N', 2, '2019-11-19 14:07:15', '2019-11-19 14:07:15');


CREATE TABLE `product` (
  `id` int(5) NOT NULL,
  `category_id` int(2) NOT NULL,
  `code` varchar(50) NOT NULL,
  `topic` varchar(200) NOT NULL,
  `detail` text NOT NULL,
  `base_price` double(8,2) NOT NULL,
  `price` double(8,2) NOT NULL,
  `stock_type` varchar(1) NOT NULL DEFAULT 'N',
  `stock` int(4) NOT NULL,
  `status` varchar(1) NOT NULL DEFAULT 'A',
  `position` int(3) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `category_id`, `code`, `topic`, `detail`, `base_price`, `price`, `stock_type`, `stock`, `status`, `position`, `created_at`, `updated_at`) VALUES
(1, 1, 'AW11', 'black tea', '', 990.00, 500.00, 'N', 0, 'A', 1, '2019-11-18 11:45:29', '2019-11-23 10:04:34'),
(2, 1, 'AW22', 'black tea with milk', '', 990.00, 500.00, 'N', 0, 'A', 1, '2019-11-18 11:46:05', '2019-11-18 18:27:11');
