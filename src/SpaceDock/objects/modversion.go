/*
 SpaceDock Backend
 API Backend for the SpaceDock infrastructure to host modfiles for various games

 SpaceDock Backend is licensed under the Terms of the MIT License.
 Copyright (c) 2017 Dorian Stoll (ThomasKerman/TMSP), RockyTV
 */

package objects

import (
    "SpaceDock"
    "os"
    "path/filepath"
)

type ModVersion struct {
    Model

    Mod             Mod `json:"-" spacedock:"lock"`
    ModID           uint `json:"mod" spacedock:"lock"`
    FriendlyVersion string `gorm:"size:64;" json:"friendly_version" spacedock:"lock"`
    Beta            bool
    GameVersion     GameVersion `json:"-" spacedock:"lock"`
    GameVersionID   uint `json:"gameversion" spacedock:"lock"`
    DownloadPath    string `gorm:"size:512" json:"download_path" spacedock:"lock"`
    Changelog       string `gorm:"size:10000"`
    SortIndex       int `json:"sort_index" spacedock:"lock"`
    FileSize        int64 `json:"file_size" spacedock:"lock"`
}

func (s *ModVersion) AfterFind() {
    if SpaceDock.DBRecursion == 2 {
        return
    }
    isRoot := SpaceDock.DBRecursion == 0
    SpaceDock.DBRecursion += 1
    SpaceDock.Database.Model(s).Related(&(s.Mod), "Mod")
    SpaceDock.Database.Model(s).Related(&(s.GameVersion), "GameVersion")
    if isRoot {
        SpaceDock.DBRecursion = 0
    }
}

func NewModVersion(mod Mod, friendly_version string, gameversion GameVersion, download_path string, beta bool) *ModVersion {
    mv := &ModVersion{
        Mod: mod,
        ModID: mod.ID,
        FriendlyVersion: friendly_version,
        Beta: beta,
        GameVersion: gameversion,
        GameVersionID: gameversion.ID,
        DownloadPath: download_path,
        Changelog: "",
        SortIndex: 0,
        FileSize: 0,
    }
    mv.Meta = "{}"
    if mv.DownloadPath != "" {
        f, err := os.Open(filepath.Join(SpaceDock.Settings.Storage, mv.DownloadPath))
        defer f.Close()
        if err == nil {
            info, err := f.Stat()
            if err == nil {
                mv.FileSize = info.Size()
            }
        }
    }
    return mv
}