

export const KATEGORI = (disabled) => {
    return {
        text: "Kategori",
        disabled: disabled,
        href: "/admin/kategori",
        icon: "fas fa-cookie-bite",
        menu: 1,
    }
};

export const UNIT_KERJA = (disabled) => {
    return {
        text: "Unit Kerja",
        disabled: disabled,
        href: "/admin/unit_kerja",
        icon: "fas fa-building",
        menu: 2,
    }
};

export const METRIC = (disabled) => {
    return {
        text: "Satuan Metrik",
        disabled: disabled,
        href: "/admin/satuan_metrik",
        icon: "fas fa-sort-numeric-down",
        menu: 3,
    }
};

export const SURAT_MASUK = (disabled) => {
    return {
        text: "Surat Masuk",
        disabled: disabled,
        href: "/admin/surat_masuk",
        icon: "fas fa-envelope",
        menu: 4,
    }
};

export const SURAT_KELUAR = (disabled) => {
    return {
        text: "Surat Keluar",
        disabled: disabled,
        href: "/admin/surat_keluar",
        icon: "fas fa-envelope-open-text",
        menu: 5,
    }
};

export const JUKNIS = (disabled) => {
    return {
        text: "Juknis",
        disabled: disabled,
        href: "/admin/juknis",
        icon: "fas fa-mail-bulk",
        menu: 6,
    }
};

export const BARANG = (disabled) => {
    return {
        text: "Inventory",
        disabled: disabled,
        href: "/admin/inventory",
        icon: "fas fa-shopping-basket",
        menu: 7,
    }
};

export const ASSET_TETAP = (disabled) => {
    return {
        text: "Inventory Tetap",
        disabled: disabled,
        href: "/admin/asset_tetap",
        icon: "fas fa-warehouse",
        menu: 8,
    }
};

export const INVENTORY_MASUK = (disabled) => {
    return {
        text: "Inventory Masuk",
        disabled: disabled,
        href: "/admin/inventory_masuk",
        icon: "fas fa-database",
        menu: 9,
    }
};

export const INVENTORY_KELUAR = (disabled) => {
    return {
        text: "Inventory Keluar",
        disabled: disabled,
        href: "/admin/inventory_keluar",
        icon: "fas fa-table",
        menu: 10,
    }
};

export const USER = (disabled) => {
    return {
        text: "User",
        disabled: disabled,
        href: "/admin/user",
        icon: "fas fa-users",
        menu: 11,
    }
};



export const DASHBOARD = {
    text: "Dashboard",
    disabled: false,
    href: "/admin",
    icon: "fas fa-home",
    menu: 0,
}

export const ADD = {
    text: "Tambah",
    disabled: true,
    href: "/admin/berkas/tambah",
    icon: "fas fa-plus",
    menu: 2,
}


export const EDIT = {
    text: "Edit",
    disabled: true,
    href: "/admin/berkas/edit",
    icon: "fas fa-edit",
    menu: 2,
}

export const INFO = {
    text: "Info",
    disabled: true,
    href: "/admin/berkas/info",
    icon: "fas fa-info",
    menu: 2,
}

export const ACCOUNT = {
    text: "Ubah Account",
    disabled: true,
    href: "/admin/user/tambah",
    icon: "fas fa-user-cog",
    menu: 1,
}

export const SETTINGS = {
    text: "Settings",
    disabled: true,
    href: "/admin/user/tambah",
    icon: "fas fa-cogs",
    menu: 1,
}

export const PASSWORD = {
    text: "Ubah Password",
    disabled: true,
    href: "/admin/user/tambah",
    icon: "fas fa-user-lock",
    menu: 1,
}