package com.example.off_chain_oracle_client.di

import com.example.off_chain_oracle_client.repository.MainRepository
import com.example.off_chain_oracle_client.network.NetworkApi
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ApplicationComponent
import javax.inject.Singleton

@InstallIn(ApplicationComponent::class)
@Module
object RepositoryModule {
    @Singleton
    @Provides
    fun provideMainRepository(
        retrofit: NetworkApi
    ): MainRepository {
        return MainRepository(
            retrofit
        )
    }
}