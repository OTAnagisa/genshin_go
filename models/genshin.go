package models

type GenshinBaseModel struct {
	/* 原神ベースモデル */

	BaseModel
	// 原神で使用されているID
	GenshinId string `json:"genshin_id"`
}

type GenshinName struct {
	/* 原神 表示名(キャラクター名, 武器名 など) */
	GenshinBaseModel
	Name     string     `json:"name"`
	Language []Language `gorm:"foreignKey:ID"`
}

func (GenshinName) TableName() string {
	return "genshin_name"
}

type GenshinIconType struct {
	/* 原神アイコン種別 */
	GenshinBaseModel
	Name string `json:"name"`
	Code int    `json:"code"`
}

func (GenshinIconType) TableName() string {
	return "genshin_icon_type"
}

type GenshinIcon struct {
	/* 原神アイコン */
	GenshinBaseModel
	Name string            `json:"name"`
	Url  string            `json:"url"`
	Type []GenshinIconType `gorm:"foreignKey:ID"`
}

func (GenshinIcon) TableName() string {
	return "genshin_icon"
}

type GenshinElement struct {
	/* 原神属性 */

	GenshinBaseModel
	Name string        `json:"name"`
	Code int           `json:"code"`
	Icon []GenshinIcon `gorm:"foreignKey:ID"`
}

func (GenshinElement) TableName() string {
	return "genshin_element"
}

type GenshinWeponType struct {
	/* 原神武器種別 */
	GenshinBaseModel
	Name string `json:"name"`
	Code int    `json:"code"`
}

func (GenshinWeponType) TableName() string {
	return "genshin_wepon_type"
}

type GenshinCharacter struct {
	/* 原神キャラクター */
	GenshinBaseModel
	DisplayName []GenshinName      `gorm:"foreignKey:ID"`
	Element     []GenshinElement   `gorm:"foreignKey:ID"`
	WeaponType  []GenshinWeponType `gorm:"foreignKey:ID"`
}

func (GenshinCharacter) TableName() string {
	return "genshin_character"
}

type GenshinCharacterIcon struct {
	/* 原神キャラクターアイコン紐づけ */
	GenshinBaseModel
	Caracter []GenshinCharacter `gorm:"foreignKey:ID"`
	Icon     []GenshinIcon      `gorm:"foreignKey:ID"`
}

func (GenshinCharacterIcon) TableName() string {
	return "genshin_character_icon"
}
