export const USER = (disabled) => {
    return {
        text: "User",
        disabled: disabled,
        href: "/admin/user",
        icon: "fas fa-users",
        menu: 6,
    }
};

export const KATEGORI = (disabled) => {
    return {
        text: "Kategori",
        disabled: disabled,
        href: "/admin/kategori",
        icon: "fas fa-cookie-bite",
        menu: 1,
    }
};

export const SURAT_MASUK = (disabled) => {
    return {
        text: "Surat Masuk",
        disabled: disabled,
        href: "/admin/surat_masuk",
        icon: "fas fa-envelope",
        menu: 2,
    }
};

export const SURAT_KELUAR = (disabled) => {
    return {
        text: "Surat Keluar",
        disabled: disabled,
        href: "/admin/surat_keluar",
        icon: "fas fa-envelope-open-text",
        menu: 3,
    }
};

export const JUKNIS = (disabled) => {
    return {
        text: "Juknis",
        disabled: disabled,
        href: "/admin/juknis",
        icon: "fas fa-mail-bulk",
        menu: 4,
    }
};

export const INVENTORY = (disabled) => {
    return {
        text: "INVENTORY",
        disabled: disabled,
        href: "/admin/inventory",
        icon: "fas fa-warehouse",
        menu: 5,
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