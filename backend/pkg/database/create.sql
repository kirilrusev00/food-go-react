SET time_zone = "+00:00";

CREATE DATABASE IF NOT EXISTS food DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE food;

CREATE TABLE foods (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `fdcId` int NOT NULL UNIQUE,
  `description` text,
  `gtinUpc` varchar(20),
  `ingredients` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO foods (fdcId, description, gtinUpc, ingredients) VALUES
  (2041155, "RAFFAELLO, ALMOND COCONUT TREAT", "009800146130", "VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT.")
