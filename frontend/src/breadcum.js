export const PROBLEM = (disabled) => {
    return {
        text: "Problem",
        disabled: disabled,
        href: "/admin/problem",
        icon: "fas fa-exclamation-triangle",
        menu: 5,
    }
};

export const PETUGAS_GAMBAR = (disabled) => {
    return {
        text: "Petugas Gambar",
        disabled: disabled,
        href: "/admin/petugas_gambar",
        icon: "fas fa-chalkboard-teacher",
        menu: 4,
    }
};

export const PETUGAS_UKUR = (disabled) => {
    return {
        text: "Petugas Ukur",
        disabled: disabled,
        href: "/admin/petugas_ukur",
        icon: "fas fa-person-booth",
        menu: 3,
    }
};

export const BERKAS = (disabled) => {
    return {
        text: "Berkas",
        disabled: disabled,
        href: "/admin/berkas",
        icon: "fas fa-id-card-alt",
        menu: 2,
    }
};

export const USER = (disabled) => {
    return {
        text: "User",
        disabled: disabled,
        href: "/admin/user",
        icon: "fas fa-users",
        menu: 1,
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