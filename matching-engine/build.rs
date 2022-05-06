fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure().build_server(true).compile(
        &[
            "../api/matching-engine/matchingEngine.proto",
            "../api/user-service/userService.proto",
        ],
        &["../api/"],
    )?;
    Ok(())
}
