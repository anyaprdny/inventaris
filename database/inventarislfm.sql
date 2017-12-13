-- MySQL dump 10.13  Distrib 5.7.20, for Linux (x86_64)
--
-- Host: localhost    Database: inventaris
-- ------------------------------------------------------
-- Server version	5.7.20-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `barang`
--

DROP TABLE IF EXISTS `barang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `barang` (
  `Kode_Barang` varchar(10) DEFAULT NULL,
  `Nama_Barang` varchar(50) DEFAULT NULL,
  `Kondisi` varchar(50) DEFAULT NULL,
  `Lokasi` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `barang`
--

LOCK TABLES `barang` WRITE;
/*!40000 ALTER TABLE `barang` DISABLE KEYS */;
INSERT INTO `barang` VALUES ('RT-01','Dispenser','Baik','Ruang Santai'),('RT-02','Kursi Komputer','Baik','Ruang Santai'),('RT-03','Troli','Baik','Selasar LFM'),('RT-04','Rak Sepatu Kayu','Baik','Selasar LFM'),('RT-05','Kasur','Baik','Ruang Santai'),('RT-06','Selimut','Baik','Ruang Santai'),('RT-07','Kotak P3K','Baik','Ruang Santai'),('RT-08','Lemari Kayu','Baik','Ruang Santai'),('RT-09','Lemari Orange','Baik','Selasar LFM'),('RT-10','Koper Perkakas','Baik','Lemari Orange'),('BK-01','Instalasi Popcorn','Baik','Selasar LFM'),('BK-02','Jack Stereo','Baik','Loker Kiri Belakang (Ruang Teknik)'),('BK-03','Kabel Audio Jack Jack','Baik','Lemari Plastik Kiri Belakang (Ruang Teknik)'),('BK-04','Kabel Microphone','Baik','Lemari Plastik Kiri Belakang (Ruang Teknik)'),('BK-05','Kabel Sound Pemutaran','Baik','Peti Hitam (Ruang Teknik)'),('BK-06','Kabel Sound 9009','Baik','Ruang 9009'),('BK-07','Proyektor','Baik','Lemari Ruang Alat'),('BK-08','Speaker','Baik','Ruang 9009'),('BK-09','Amplifier','Baik','Meja Kiri Depan (Ruang Teknik)'),('BK-10','Stereo Power Mixer','Cacat','Meja Kiri Depan (Ruang Teknik)'),('BK-11','Terminal A','Cacat','Lemari Plastik Kiri Belakang (Ruang Teknik)'),('BK-12','Terminal B','Baik','Lemari Plastik Kiri Belakang (Ruang Teknik)'),('DU-01','Wireless Flash Trigger','Rusak','Toolbox Ruang Alat'),('DU-02','Card Reader A','Baik','Loker Danus'),('DU-03','Card Reader B','Baik','Ruang Teknik'),('DU-04','Card Reader C','Rusak','Ruang Teknik'),('DU-05','SD Card Transcend 16GB','Baik','Loker Danus'),('DU-06','SD Card SanDisk 8GB','Baik','Loker Danus'),('DU-07','Hard Disk Wisuda Oktober','Baik','Loker Danus'),('DU-08','Hard Disk Wisuda April','Baik','Loker Danus'),('DU-09','Hard Disk Wisuda Juli','Baik','Loker Danus'),('DU-10','Wireless Router TPLINK 300','Baik','Loker Danus'),('DU-11','Hard Disk Proyek','Baik','Loker Danus'),('DU-12','Peralatan Marketing','Baik','Ruang Alat'),('DU-13','Handycam JVC GZ-HM860','Baik','Ruang Alat'),('DU-14','Printer HP Laser JET P1102','Baik','Ruang Alat'),('DU-15','Cash Box','Baik','Ruang Alat'),('DU-16','Mesin EDC','Baik','Loker Danus'),('DS-01','Kamera Canon 700D','Baik','Lemari Doksos'),('DS-02','Lensa 18-55mm Canon','Baik','Lemari Doksos'),('DS-03','Charger LC-E8 Canon','Baik','Lemari Doksos'),('DS-04','SD Card','Baik','Lemari Doksos'),('DS-05','Kit Pembersih Kamera','Baik','Lemari Doksos'),('DS-06','Hard Disk Doksos','Baik','Loker Manajer Doksos'),('DS-07','External Flash','Baik','Lemari Doksos'),('VD-01','Tripod Velbon EX383','Baik','Kotak Tripod Ruang Alat'),('VD-02','Dolly','Baik','Ruang Alat'),('VD-03','Zoom H4N','Baik','Ruang Alat'),('VD-04','Halogen','Baik','Ruang Alat'),('VD-05','Clapperboard','Baik','Ruang Alat'),('VD-06','Green Screen','Baik','Ruang Alat'),('VD-07','Slider','Baik','Ruang Alat'),('VD-08','Mic Boom','Baik','Ruang Alat'),('VD-09','Windshield','Baik','Ruang Alat'),('VD-10','Boom Pole','Baik','Ruang Alat'),('VD-11','Monopod','Baik','Ruang Alat'),('VD-12','Glide Cam HD200','Baik','Ruang Alat'),('FT-01','Lampu Lighting','Rusak','Ruang Alat'),('FT-02','Softbox','Baik','Ruang Alat'),('FT-03','Lightstand','Baik','Ruang Alat'),('FT-04','Developing Film Kit','Baik','Ruang Alat'),('FT-05','Lensa Tamron 85-210mm f/4','Baik','Ruang Alat'),('FT-06','Schneider-Kreuznach 50mm','Baik','Ruang Alat'),('KN-01','Katalog Festival JAFF 2016','Baik','Rak Buku Ruang Santai'),('KN-02','Buku \'Mau Dibawa Kemana Sinema Kita?\'','Baik','Rak Buku Ruang Santai'),('KN-03','Katalog Ganesha Film Festival 2014','Baik','Rak Buku Ruang Santai'),('KN-04','Katalog Ganesha Film Festival 2016','Baik','Rak Buku Ruang Santai'),('KN-05','Buku \'Film Production Management\'','Baik','Rak Buku Ruang Santai'),('KN-06','DVD 3 Hari untuk Selamanya','Baik','Di Atas Televisi Ruang Santai'),('KN-07','DVD Janji Joni','Baik','Di Atas Televisi Ruang Santai'),('KN-08','DVD Selamat Pagi, Malam','Baik','Di Atas Televisi Ruang Santai'),('KN-09','DVD Realita Cinta dan Rock n Roll','Baik','Di Atas Televisi Ruang Santai'),('KN-10','DVD The Grand Budapest Hotel','Baik','Di Atas Televisi Ruang Santai'),('KN-11','DVD Kompilasi 2017','baik','Di Atas TV Ruang Santai');
/*!40000 ALTER TABLE `barang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bidang`
--

DROP TABLE IF EXISTS `bidang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bidang` (
  `ID_Bidang` char(2) DEFAULT NULL,
  `Nama_Bidang` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bidang`
--

LOCK TABLES `bidang` WRITE;
/*!40000 ALTER TABLE `bidang` DISABLE KEYS */;
INSERT INTO `bidang` VALUES ('BK','Bioskop Kampus'),('DS','Dokumentasi Sosial'),('DU','Dana Usaha'),('FT','Fotografi'),('KN','Kineklub'),('RT','Biro Rumah Tangga'),('VD','Videografi');
/*!40000 ALTER TABLE `bidang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventarisbidang`
--

DROP TABLE IF EXISTS `inventarisbidang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventarisbidang` (
  `BidangID` char(2) NOT NULL,
  `BarangKode` varchar(10) NOT NULL,
  PRIMARY KEY (`BidangID`,`BarangKode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventarisbidang`
--

LOCK TABLES `inventarisbidang` WRITE;
/*!40000 ALTER TABLE `inventarisbidang` DISABLE KEYS */;
INSERT INTO `inventarisbidang` VALUES ('BK','BK-01'),('BK','BK-02'),('BK','BK-03'),('BK','BK-04'),('BK','BK-05'),('BK','BK-06'),('BK','BK-07'),('BK','BK-08'),('BK','BK-09'),('BK','BK-10'),('BK','BK-11'),('BK','BK-12'),('DS','DS-01'),('DS','DS-02'),('DS','DS-03'),('DS','DS-04'),('DS','DS-05'),('DS','DS-06'),('DS','DS-07'),('DU','DU-01'),('DU','DU-02'),('DU','DU-03'),('DU','DU-04'),('DU','DU-05'),('DU','DU-06'),('DU','DU-07'),('DU','DU-08'),('DU','DU-09'),('DU','DU-10'),('DU','DU-11'),('DU','DU-12'),('DU','DU-13'),('DU','DU-14'),('DU','DU-15'),('DU','DU-16'),('FT','FT-01'),('FT','FT-02'),('FT','FT-03'),('FT','FT-04'),('FT','FT-05'),('FT','FT-06'),('KN','KN-01'),('KN','KN-02'),('KN','KN-03'),('KN','KN-04'),('KN','KN-05'),('KN','KN-06'),('KN','KN-07'),('KN','KN-08'),('KN','KN-09'),('KN','KN-10'),('RT','RT-01'),('RT','RT-02'),('RT','RT-03'),('RT','RT-04'),('RT','RT-05'),('RT','RT-06'),('RT','RT-07'),('RT','RT-08'),('RT','RT-09'),('RT','RT-10'),('VD','VD-01'),('VD','VD-02'),('VD','VD-03'),('VD','VD-04'),('VD','VD-05'),('VD','VD-06'),('VD','VD-07'),('VD','VD-08'),('VD','VD-09'),('VD','VD-10'),('VD','VD-11'),('VD','VD-12');
/*!40000 ALTER TABLE `inventarisbidang` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-13 12:50:04
