
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id int AUTO_INCREMENT PRIMARY KEY,
    telephone varchar(50) NOT NULL,
    password varchar(200) NOT NULL,
    name varchar(20) NOT NULL,
    sex varchar(20) DEFAULT NULL,
    role varchar(20) NOT NULL,
    UNIQUE KEY (`telephone`)
);

-- 申请信息，包括维修和适驾（编号、用户编号、车辆编号、服务类型、问题描述、空闲时间、申请是否完成）
DROP TABLE IF EXISTS applications;
CREATE TABLE applications
(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    car_id int NOT NULL,
    class varchar(20) NOT NULL,
    description varchar(250) NOT NULL,
    time varchar(100) NOT NULL,
    is_end TINYINT(1) NOT NULL
);

-- 汽车信息（汽车编号，品牌，车型名称，图片，厂商指导价，颜色、长宽高、发动机、变速箱、最大马力、最大功率、驱动方式、燃料形式）

DROP TABLE IF EXISTS cars;
CREATE TABLE cars
(
    id int AUTO_INCREMENT PRIMARY KEY,
    brand varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    image varchar(250) DEFAULT NULL,
    price varchar(50) DEFAULT NULL,
    color varchar(20) DEFAULT NULL,
    size varchar(50) DEFAULT NULL,
    motor varchar(50) DEFAULT NULL,
    gearbox varchar(50) DEFAULT NULL,
    max_horse varchar(50) DEFAULT NULL,
    max_power varchar(50) DEFAULT NULL,
    drive_mode varchar(50) DEFAULT NULL,
    fuel varchar(50) DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS collections;
CREATE TABLE collections
(
    id int NOT NULL PRIMARY KEY,
    user_id int NOT NULL,
    car_id int NOT NULL
#     FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
#     FOREIGN KEY (`car_id`) REFERENCES `cars` (`id`)
);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders
(
    id int NOT NULL PRIMARY KEY,
    user_id int NOT NULL,
    car_id int NOT NULL
#     FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
#     FOREIGN KEY (`car_id`) REFERENCES `cars` (`id`)
);

INSERT INTO cars VALUES(1, '奥迪', '奥迪Q3 2015款 30 TFSI 标准型', '‪‪./image/aodi1.jpg', '24.98万', '黑色', '4385*1831*1589', '1.4T 150马力 L4', '6挡手动', '150', '110', '前置前驱', '汽油');
INSERT INTO cars VALUES(2, '奥迪', '奥迪Q3 2016款 30 TFSI 典藏版 智领型', '‪‪./image/aodi2.jpg', '25.18万', '棕色', '4385*1831*1589', '1.4T 150马力 L4', '6挡双离合', '150', '110', '前置前驱', '汽油');
INSERT INTO cars VALUES(3, '奥迪', '奥迪Q3 2015款 35 TFSI quattro 豪华型', './image/aodi3.png', '39.68万', '白色', '4385*1831*1589', '2.0T 170马力 L4', '7挡双离合', '170', '125', '前置四驱', '汽油');
INSERT INTO cars VALUES(4, '奥迪', '奥迪Q3 2015款 40 TFSI quattro 豪华型', './image/aodi4.jpg', '42.88万', '松露米色', '4385*1831*1589', '2.0T 200马力 L4', '7挡双离合', '200', '147', '前置四驱', '汽油');
INSERT INTO cars VALUES(5, '奥迪', '奥迪Q3 2013款 35 TFSI 进取型', './image/aodi5.png','28.50万', '白色', '4385*1831*1608', '2.0T 170马力 L4', '7挡双离合', '170', '125', '前置前驱','汽油');

INSERT INTO cars VALUES(6, '宝马', '宝马3系 2016款 316i 时尚型', '‪./image/baoma1.jpg', '28.30万', '白色', '4643*1811*1454', '1.6T 136马力 L4', '8挡手自一体', '136', '100', '前置后驱', '汽油');
INSERT INTO cars VALUES(7, '宝马', '宝马3系 2015款 320Li 豪华设计套装', './image/baoma2.jpg', '39.68万', '白色', '4734*1811*1453', '2.0T 184马力 L4', '8挡手自一体', '184', '135', '前置后驱', '汽油');
INSERT INTO cars VALUES(8, '宝马', '宝马3系 2015款 320i 进取型', './image/baoma3.jpg', '32.00万', '威尼托米色', '4624*1811*1455', '2.0T 184马力 L4', '8挡手自一体', '184', '135', '前置后驱', '汽油');
INSERT INTO cars VALUES(9, '宝马', '宝马3系 2014款 320Li 手动型', './image/baoma4.jpg', '32.98万', '白色', '4734*1811*1453', '2.0T 184马力 L4', '6挡手动', '184', '135', '前置后驱', '汽油');
INSERT INTO cars VALUES(10, '宝马', '宝马X6 2015款 xDrive35i 豪华型', './image/baoma5.jpg', '99.80万', '黑色', '4929*1989*1709', '3.0T 306马力 L6', '8挡手自一体', '306', '225','前置四驱', '汽油');

INSERT INTO cars VALUES(11, '保时捷', 'Boxster 2013款 Boxster 2.7L', './image/baoshijie1.jpg', '72.60万', '黑色', '4374*1801*1282', '2.7L 265马力 H6', '7挡双离合', '265', '195', '中置后驱', '汽油');
INSERT INTO cars VALUES(12, '保时捷', 'Boxster 2015款 Boxster Style Edition 2.7L', './image/baoshijie2.jpg', '74.70万', '蓝色', '4374*1801*1282', '2.7L 265马力 H6', '7挡双离合', '265', '195', '中置后驱', '汽油');
INSERT INTO cars VALUES(13, '保时捷', 'Boxster 2013款 Boxster S 3.4L', './image/baoshijie3.jpg', '97.00万', '白色', '4374*1801*1282', '3.4L 316马力 H6', '7挡双离合', '316', '232', '中置后驱', '汽油');
INSERT INTO cars VALUES(14, '保时捷', 'Boxster 2014款 Boxster GTS 3.4L', './image/baoshijie4.jpg', '106.50万', '珀金灰', '4404*1801*1273', '3.4L 330马力 H6', '7挡双离合', '330', '243', '中置后驱', '汽油');
INSERT INTO cars VALUES(15, '保时捷', 'Cayman 2014款 Cayman GTS 3.4L', './image/baoshijie5.jpg', '111.80万', '红色', '4404*1801*1284', '3.4L 340马力 H6', '7挡双离合', '340', '250', '中置后驱', '汽油');

INSERT INTO cars VALUES(16, '奔驰', '奔驰C级 2015款 C 300 L', './image/benchi1.jpg', '48.90万', '白色', '4783*1810*1442', '2.0T 245马力 L4', '7挡手自一体', '245', '180', '前置后驱', '汽油');
INSERT INTO cars VALUES(17, '奔驰', '奔驰C级 2015款 C 300 L 运动型', './image/benchi2.jpg', '48.90万', '灰色', '4783*1810*1442', '2.0T 245马力 L4', '7挡手自一体', '245', '180', '前置后驱', '汽油');
INSERT INTO cars VALUES(18, '奔驰', '奔驰C级 2016款 C 350 eL', './image/benchi3.jpeg', '59.90万', '蓝色', '4771*1810*1442', '2.0T 211马力 L4', '7挡手自一体', '211', '155', '前置后驱', '插电式混合动力');
INSERT INTO cars VALUES(19, '奔驰', '奔驰C级 2015款 C 200 运动版', './image/benchi4.jpeg', '31.48万', '灰色','4714*1810*1442', '2.0T 184马力 L4', '7挡手自一体', '184', '135', '前置后驱', '汽油');
INSERT INTO cars VALUES(20, '奔驰', '奔驰C级 2008款 C 200K 标准型', './image/benchi5.jpg', '34.80万', '银色', '4581*1770*1448', '1.8T 184马力 L4', '5挡手自一体', '184', '135', '前置后驱', '汽油');

INSERT INTO cars VALUES(21, '长安', '逸动 2015款 纯电动尊贵型', './image/changan1.jpg','24.99万', '白色', '4620*1820*1515', '纯电动 122马力', '电动车单速变速箱', '122', '90', '前置前驱', '纯电动');
INSERT INTO cars VALUES(22, '长安', '逸动 2014款 1.6L 手动精英型', './image/changan2.jpg', '7.49万', '灰色', '4620*1820*1490', '1.6L 125马力 L4', '5挡手动', '125', '92', '前置前驱', '汽油');
INSERT INTO cars VALUES(23, '长安', '逸动 2014款 1.6L 自动精英型', './image/changan3.jpg', '8.39万', '灰色', '4620*1820*1490', '1.6L 125马力 L4', '4挡手自一体', '125', '92', '前置前驱', '汽油');
INSERT INTO cars VALUES(24, '长安', '逸动 2014款 1.6L 手动豪华型', './image/changan4.jpeg', '7.99万', '白色', '4620*1820*1490', '1.6L 125马力 L4', '5挡手动', '125', '92', '前置前驱', '汽油');
INSERT INTO cars VALUES(25, '长安', '悦翔V7 2015款 1.6L 手动乐动型 国IV', './image/changan5.jpg', '6.59万', '黑色', '4530*1745*1498', '1.6L 124马力 L4', '5挡手动', '124', '91', '前置前驱', '汽油');

INSERT INTO cars VALUES(26, '大众', 'POLO 2015款 1.4TSI GTI', './image/dazong1.jpg', '14.69万', '白色', '3975*1682*1487', '1.4T 150马力 L4', '7挡双离合', '150', '110', '前置前驱', '汽油');
INSERT INTO cars VALUES(27, '大众', 'POLO 2014款 1.4L 手动舒适版', './image/dazong2.jpg', '8.69万', '蓝色', '3970*1682*1462', '1.4L 90马力 L4', '5挡手动', '90', '66', '前置前驱', '汽油');
INSERT INTO cars VALUES(28, '大众', 'POLO 2014款 1.4L 自动舒适版', './image/dazong3.jpg', '9.89万', '白色', '3970*1682*1462', '1.4L 90马力 L4', '6挡手自一体', '90', '66', '前置前驱', '汽油');
INSERT INTO cars VALUES(29, '大众', '途观 2016款 330TSI 自动四驱旗舰版', './image/dazong4.jpg', '31.58万', '黑色', '4506*1809*1685', '2.0T 200马力 L4', '6挡手自一体', '200', '147', '前置四驱', '汽油');
INSERT INTO cars VALUES(30, '大众', '途观 2015款 1.4TSI 手动两驱蓝驱版', './image/dazong5.jpg', '19.98万', '蓝色', '4506*1809*1685', '1.4T 131马力 L4', '6挡手动', '131', '96', '前置前驱', '汽油');

INSERT INTO cars VALUES(31, '东风', '帅客 2016款 1.6L 手动舒适型', './image/dongfeng1.jpg', '8.68万', '白色', '4500*1695*1825', '1.6L 113马力 L4', '5挡手动', '113', '83', '前置前驱', '汽油');
INSERT INTO cars VALUES(32, '东风', '帅客 2012款 2.0L 自动旗舰型7座', './image/dongfeng2.jpg', '11.58万', '灰色', '4420*1695*1825', '2.0L 147马力 L4', '4挡自动', '147', '108', '前置前驱', '汽油');
INSERT INTO cars VALUES(33, '东风', '帅客 2014款 1.5L 手动标准型5座 国IV', './image/dongfeng3.jpg', '6.48万', '橙红色', '4420*1695*1825', '1.5L 102马力 L4', '5挡手动', '102', '75', '前置前驱', '汽油');
INSERT INTO cars VALUES(34, '东风', '帅客 2014款 1.5L 手动标准型7座 国IV', './image/dongfeng4.jpg', '6.78万', '银色', '4420*1695*1825', '1.5L 102马力 L4', '5挡手动', '102', '75', '前置前驱', '汽油');
INSERT INTO cars VALUES(35, '东风', '帅客 2013款 改款 1.5L 手动舒适型7座 国IV', './image/dongfeng5.jpg', '7.18万', '白色', '4420*1695*1825', '1.5L 102马力 L4', '5挡手动', '102', '75', '前置前驱', '汽油');

INSERT INTO cars VALUES(36, '福特', '锐界(进口) 2011款 3.5L 精锐型', './image/fute1.jpg', '36.98万', '黑色', '4699*1930*1705', '3.5L 288马力 V6', '6挡手自一体', '288', '212', '前置四驱', '汽油');
INSERT INTO cars VALUES(37, '福特', '锐界(进口) 2011款 3.5L 精锐天窗版', './image/fute2.jpg', '38.98万', '黑色', '4699*1930*1705', '3.5L 288马力 V6', '6挡手自一体', '288', '212', '前置四驱', '汽油');
INSERT INTO cars VALUES(38, '福特', '锐界(进口) 2011款 3.5L 尊锐型', './image/fute3.jpg', '42.98万', '白色', '4699*1930*1705', '3.5L 288马力 V6', '6挡手自一体', '288', '212', '前置四驱', '汽油');
INSERT INTO cars VALUES(39, '福特', '锐界(进口) 2009款 3.5 V6', './image/fute4.jpg', '53.80万', '橙色', '4717*1955*1702', '3.5L 269马力 V6', '6挡自动', '269', '198', '前置四驱', '汽油');
INSERT INTO cars VALUES(40, '福特', '探险者 2016款 2.3T 风尚版', './image/fute5.jpg', '44.98万', '墨绿色', '5037*2005*1816', '2.3T 276马力 L4', '6挡手自一体', '276', '203', '前置四驱', '汽油');

INSERT INTO cars VALUES(41, 'Jeep', '牧马人 2012款 3.6L 四门版 Rubicon', './image/Jeep1.jpg', '54.99万', '白色', '4751*1877*1840', '3.6L 284马力 V6', '5挡手自一体', '284', '209', '前置四驱', '汽油');
INSERT INTO cars VALUES(42, 'Jeep', '牧马人 2011款 3.8L 两门版 Sahara', './image/Jeep2.jpg', '41.89万', '黑色', '4223*1873*1865', '3.8L 199马力 V6', '4挡自动', '199', '146', '前置四驱', '汽油');
INSERT INTO cars VALUES(43, 'Jeep', '指南者 2015款 2.0L 两驱运动版', './image/Jeep3.jpg', '22.19万', '白色', '4465*1812*1663', '2.0L 156马力 L4', '6挡手自一体', '156', '115', '前置前驱', '汽油');
INSERT INTO cars VALUES(44, 'Jeep', '指南者 2015款 2.0L 两驱豪华版', './image/Jeep4.jpg', '24.49万', '黑色', '4465*1812*1663', '2.0L 156马力 L4', '6挡手自一体', '156', '115', '前置前驱', '汽油');
INSERT INTO cars VALUES(45, 'Jeep', '指南者 2014款 改款 2.4L 四驱舒适版', './image/Jeep5.jpg', '24.59万', '黑色', '4465*1812*1663', '2.4L 170马力 L4', '6挡手自一体', '170', '125', '前置四驱', '汽油');

INSERT INTO cars VALUES(46, '路虎', '揽胜运动版 2016款 3.0 SD V6 Hybrid HSE Dynamic', './image/luhu1.jpg', '119.80万', '白色', '4850*2073*1780', '3.0T 292马力 V6', '8挡手自一体', '292', '215', '前置四驱', '油电混合');
INSERT INTO cars VALUES(47, '路虎', '揽胜运动版 2016款 5.0 V8 SVR', './image/luhu2.jpg', '229.80万', '象牙白', '4850*2073*1780', '5.0T 551马力 V8', '8挡手自一体', '551', '405', '前置四驱', '汽油');
INSERT INTO cars VALUES(48, '路虎', '揽胜运动版 2015款 3.0 SDV6 HSE', './image/luhu3.jpg', '98.80万', '白色', '4850*2073*1780', '3.0T 258马力 V6', '8挡手自一体', '258', '190', '前置四驱', '柴油');
INSERT INTO cars VALUES(49, '路虎', '神行者2 2015款 2.0T Si4 XS典藏版', './image/luhu4.jpg', '53.80万', '白色', '4500*1910*1740', '2.0T 240马力 L4 ', '6挡手自一体', '240', '176.5', '前置四驱', '汽油');
INSERT INTO cars VALUES(50, '路虎',  '神行者2 2015款 2.0T Si4 HSE Luxury典藏版', './image/luhu5.jpg', '60.80万', '黑色', '4500*1910*1740', '2.0T 240马力 L4 ', '6挡手自一体', '240', '176.5', '前置四驱', '汽油');

INSERT INTO cars VALUES(51, '现代', '途胜 2013款 2.0L 手动两驱舒适型', './image/xiandai1.jpg', '16.58万', '黑色', '4345*1795*1680', '2.0L 141马力 L4', '5挡手动', '141', '104', '前置前驱', '汽油');
INSERT INTO cars VALUES(52, '现代', '途胜 2013款 2.0L 手动两驱时尚型', './image/xiandai2.jpg', '16.98万', '黑色', '4345*1797*1730', '2.0L 141马力 L4', '5挡手动', '141', '104', '前置前驱', '汽油');
INSERT INTO cars VALUES(53, '现代', '途胜 2013款 2.0L 自动两驱舒适型', './image/xiandai3.jpg', '18.38万', '银色', '4345*1797*1730', '2.0L 141马力 L4', '4挡手自一体', '141', '104', '前置前驱', '汽油');
INSERT INTO cars VALUES(54, '现代', '悦动 2015款 1.6L 手动舒适型', './image/xiandai4.jpg', '9.98万', '黑色', '4543*1777*1490', '1.6L 123马力 L4', '5挡手动', '123', '90.4', '前置前驱', '汽油');
INSERT INTO cars VALUES(55, '现代', '悦动 2015款 1.6L 自动舒适型',  './image/xiandai5.jpg', '10.98万', '黑色', '4543*1777*1490', '1.6L 123马力 L4', '4挡自动', '123', '90.4', '前置前驱', '汽油');

-- 申请信息 applications
insert into applications values (1,1,1,'美容保养','镀膜','2024-1-2',0);
insert into applications values (2,3,18,'试驾预约','','2024-1-5',1);
insert into applications values (3,4,5,'美容保养','打蜡','2024-1-3',0);
insert into applications values (4,6,19,'试驾预约','','2024-1-2',1);
insert into applications values (5,10,22,'美容保养','抛光','2024-1-1',1);
insert into applications values (6,7,8,'车辆维修','修理发动机','2024-1-6',1);
insert into applications values (7,8,22,'试驾预约','','2024-1-8',0);
insert into applications values (8,4,4,'美容保养','镀膜','2024-1-2',0);
insert into applications values (9,15,27,'试驾预约','','2024-1-7',0);
insert into applications values (10,10,23,'车辆维修','修理雨刷','2024-1-3',0);
insert into applications values (11,13,34,'试驾预约','','2024-1-7',1);
insert into applications values (12,9,21,'美容保养','打蜡','2024-1-2',1);
insert into applications values (13,6,7,'车辆维修','修理车胎','2024-1-3',1);

-- 收藏 collections
INSERT INTO collections VALUES(1, 1, 16);
INSERT INTO collections VALUES(2, 1, 17);
INSERT INTO collections VALUES(3, 1, 18);
INSERT INTO collections VALUES(4, 1, 19);
INSERT INTO collections VALUES(5, 1, 20);
INSERT INTO collections VALUES(6, 2, 21);
INSERT INTO collections VALUES(7, 2, 22);
INSERT INTO collections VALUES(8, 3, 23);
INSERT INTO collections VALUES(9, 4, 24);
INSERT INTO collections VALUES(10, 5, 25);
INSERT INTO collections VALUES(11, 6, 26);
INSERT INTO collections VALUES(12, 7, 27);

-- 订单 orders
INSERT INTO orders VALUES(1, 1, 1);
INSERT INTO orders VALUES(2, 2, 2);
INSERT INTO orders VALUES(3, 2, 3);
INSERT INTO orders VALUES(4, 4, 4);
INSERT INTO orders VALUES(5, 4, 5);
INSERT INTO orders VALUES(6, 5, 6);
INSERT INTO orders VALUES(7, 6, 7);
INSERT INTO orders VALUES(8, 7, 8);
INSERT INTO orders VALUES(9, 8, 9);
INSERT INTO orders VALUES(10, 9, 20);
INSERT INTO orders VALUES(11, 9, 21);
INSERT INTO orders VALUES(12, 10, 22);
INSERT INTO orders VALUES(13, 10, 23);

