-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 23/01/2024 às 03:28
-- Versão do servidor: 10.4.28-MariaDB
-- Versão do PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `harmony`
--

-- --------------------------------------------------------

--
-- Estrutura para tabela `prompts`
--

CREATE TABLE `prompts` (
  `id` int(11) NOT NULL,
  `text` varchar(255) NOT NULL,
  `model` int(11) NOT NULL,
  `tags` int(11) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Despejando dados para a tabela `prompts`
--

INSERT INTO `prompts` (`id`, `text`, `model`, `tags`, `created_at`, `updated_at`) VALUES
(1, 'Sample1', 0, 0, 1705976805213607600, 1705976805213607600),
(2, 'Sample2', 0, 0, 1705976816294213900, 1705976816294213900);

-- --------------------------------------------------------

--
-- Estrutura para tabela `users`
--

CREATE TABLE `users` (
  `id` int(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Despejando dados para a tabela `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'harmony', '$2a$10$4MihoYX8ttS.nZhyKEskIe8xUknJv.vHT5nHwxhOCQ2adSGgQfEOS'),
(2, 'harmony2', '$2a$10$B3N4T.2fzilJlNC.1xOc..d6SGg6MfbVlZyfhpq0o0a30mpLdIlK6'),
(3, 'harmony3', '$2a$10$wodJHWRWnG5DMx9idkhbO.eJs/CjFBZWMOiwrA7I3nt59TIJdcU9i'),
(4, 'harmony4', '$2a$10$//tf4DA2SurQVg7HGe//Fu3NCOupPxH4hQnKfPkCTO0RkXt2Y4z5a'),
(5, 'harmony5', '$2a$10$Etb3BjoA5LtXEp9XKlWv3.7Sy7JE3iifnj3xIHFIhCUC.nyoXoqv2'),
(6, 'harmony10', '$2a$10$i9mQm8ZUQj.5x7ySlHxzJ./S/z6AOM0MLbLSMpUTmVyq6aJ6L0GLm');

--
-- Índices para tabelas despejadas
--

--
-- Índices de tabela `prompts`
--
ALTER TABLE `prompts`
  ADD PRIMARY KEY (`id`);

--
-- Índices de tabela `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT para tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `prompts`
--
ALTER TABLE `prompts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT de tabela `users`
--
ALTER TABLE `users`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
