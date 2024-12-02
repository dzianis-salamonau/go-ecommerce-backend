# Define the base directory for your project
$baseDir = "c:\Users\user\Documents\go\go-ecommerce-backend\"

# Define the folder structure
$folders = @(
    "$baseDir/cmd/ecommerce",
    "$baseDir/pkg/users",
    "$baseDir/pkg/products",
    "$baseDir/pkg/orders",
    "$baseDir/pkg/payments",
    "$baseDir/pkg/shipping",
    "$baseDir/internal/auth",
    "$baseDir/internal/db",
    "$baseDir/internal/services"
)

# Define the file structure
$files = @(
    "$baseDir/cmd/github.com/dzianis-salamonau/go-ecommerce-backend/main.go",
    "$baseDir/pkg/users/users.go",
    "$baseDir/pkg/products/products.go",
    "$baseDir/pkg/orders/orders.go",
    "$baseDir/pkg/payments/payments.go",
    "$baseDir/pkg/shipping/shipping.go",
    "$baseDir/internal/auth/auth.go",
    "$baseDir/internal/db/db.go",
    "$baseDir/internal/services/services.go"
)

# Create the folders
foreach ($folder in $folders) {
    if (-not (Test-Path -Path $folder)) {
        New-Item -Path $folder -ItemType Directory -Force
        Write-Host "Created folder: $folder"
    }
}

# Create the files
foreach ($file in $files) {
    if (-not (Test-Path -Path $file)) {
        New-Item -Path $file -ItemType File -Force
        Write-Host "Created file: $file"
    }
}

Write-Host "Project structure created successfully!"
